package p2p

import (
	"fmt"
	golog "log"
	"math"
	"net"
	"testing"
	"time"

	"github.com/tendermint/go-amino"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/exfury/fuxchain/libs/tendermint/crypto"
	"github.com/exfury/fuxchain/libs/tendermint/crypto/ed25519"
	"github.com/exfury/fuxchain/libs/tendermint/libs/bytes"
	"github.com/exfury/fuxchain/libs/tendermint/libs/log"

	"github.com/exfury/fuxchain/libs/tendermint/config"
	tmconn "github.com/exfury/fuxchain/libs/tendermint/p2p/conn"
)

func TestPeerBasic(t *testing.T) {
	assert, require := assert.New(t), require.New(t)

	// simulate remote peer
	rp := &remotePeer{PrivKey: ed25519.GenPrivKey(), Config: cfg}
	rp.Start()
	defer rp.Stop()

	p, err := createOutboundPeerAndPerformHandshake(rp.Addr(), cfg, tmconn.DefaultMConnConfig())
	require.Nil(err)

	err = p.Start()
	require.Nil(err)
	defer p.Stop()

	assert.True(p.IsRunning())
	assert.True(p.IsOutbound())
	assert.False(p.IsPersistent())
	p.persistent = true
	assert.True(p.IsPersistent())
	assert.Equal(rp.Addr().DialString(), p.RemoteAddr().String())
	assert.Equal(rp.ID(), p.ID())
}

func TestPeerSend(t *testing.T) {
	assert, require := assert.New(t), require.New(t)

	config := cfg

	// simulate remote peer
	rp := &remotePeer{PrivKey: ed25519.GenPrivKey(), Config: config}
	rp.Start()
	defer rp.Stop()

	p, err := createOutboundPeerAndPerformHandshake(rp.Addr(), config, tmconn.DefaultMConnConfig())
	require.Nil(err)

	err = p.Start()
	require.Nil(err)

	defer p.Stop()

	assert.True(p.CanSend(testCh))
	assert.True(p.Send(testCh, []byte("Asylum")))
}

func createOutboundPeerAndPerformHandshake(
	addr *NetAddress,
	config *config.P2PConfig,
	mConfig tmconn.MConnConfig,
) (*peer, error) {
	chDescs := []*tmconn.ChannelDescriptor{
		{ID: testCh, Priority: 1},
	}
	reactorsByCh := map[byte]Reactor{testCh: NewTestReactor(chDescs, true)}
	pk := ed25519.GenPrivKey()
	pc, err := testOutboundPeerConn(addr, config, false, pk)
	if err != nil {
		return nil, err
	}
	timeout := 1 * time.Second
	ourNodeInfo := testNodeInfo(addr.ID, "host_peer")
	peerNodeInfo, err := handshake(pc.conn, timeout, ourNodeInfo)
	if err != nil {
		return nil, err
	}

	p := newPeer(pc, mConfig, peerNodeInfo, reactorsByCh, chDescs, func(p Peer, r interface{}) {})
	p.SetLogger(log.TestingLogger().With("peer", addr))
	return p, nil
}

func testDial(addr *NetAddress, cfg *config.P2PConfig) (net.Conn, error) {
	if cfg.TestDialFail {
		return nil, fmt.Errorf("dial err (peerConfig.DialFail == true)")
	}

	conn, err := addr.DialTimeout(cfg.DialTimeout)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func testOutboundPeerConn(
	addr *NetAddress,
	config *config.P2PConfig,
	persistent bool,
	ourNodePrivKey crypto.PrivKey,
) (peerConn, error) {

	var pc peerConn
	conn, err := testDial(addr, config)
	if err != nil {
		return pc, errors.Wrap(err, "Error creating peer")
	}

	pc, err = testPeerConn(conn, config, true, persistent, ourNodePrivKey, addr)
	if err != nil {
		if cerr := conn.Close(); cerr != nil {
			return pc, errors.Wrap(err, cerr.Error())
		}
		return pc, err
	}

	// ensure dialed ID matches connection ID
	if addr.ID != pc.ID() {
		if cerr := conn.Close(); cerr != nil {
			return pc, errors.Wrap(err, cerr.Error())
		}
		return pc, ErrSwitchAuthenticationFailure{addr, pc.ID()}
	}

	return pc, nil
}

type remotePeer struct {
	PrivKey    crypto.PrivKey
	Config     *config.P2PConfig
	addr       *NetAddress
	channels   bytes.HexBytes
	listenAddr string
	listener   net.Listener
}

func (rp *remotePeer) Addr() *NetAddress {
	return rp.addr
}

func (rp *remotePeer) ID() ID {
	return PubKeyToID(rp.PrivKey.PubKey())
}

func (rp *remotePeer) Start() {
	if rp.listenAddr == "" {
		rp.listenAddr = "127.0.0.1:0"
	}

	l, e := net.Listen("tcp", rp.listenAddr) // any available address
	if e != nil {
		golog.Fatalf("net.Listen tcp :0: %+v", e)
	}
	rp.listener = l
	rp.addr = NewNetAddress(PubKeyToID(rp.PrivKey.PubKey()), l.Addr())
	if rp.channels == nil {
		rp.channels = []byte{testCh}
	}
	go rp.accept()
}

func (rp *remotePeer) Stop() {
	rp.listener.Close()
}

func (rp *remotePeer) Dial(addr *NetAddress) (net.Conn, error) {
	conn, err := addr.DialTimeout(1 * time.Second)
	if err != nil {
		return nil, err
	}
	pc, err := testInboundPeerConn(conn, rp.Config, rp.PrivKey)
	if err != nil {
		return nil, err
	}
	_, err = handshake(pc.conn, time.Second, rp.nodeInfo())
	if err != nil {
		return nil, err
	}
	return conn, err
}

func (rp *remotePeer) accept() {
	conns := []net.Conn{}

	for {
		conn, err := rp.listener.Accept()
		if err != nil {
			golog.Printf("Failed to accept conn: %+v", err)
			for _, conn := range conns {
				_ = conn.Close()
			}
			return
		}

		pc, err := testInboundPeerConn(conn, rp.Config, rp.PrivKey)
		if err != nil {
			golog.Fatalf("Failed to create a peer: %+v", err)
		}

		_, err = handshake(pc.conn, time.Second, rp.nodeInfo())
		if err != nil {
			golog.Fatalf("Failed to perform handshake: %+v", err)
		}

		conns = append(conns, conn)
	}
}

func (rp *remotePeer) nodeInfo() NodeInfo {
	return DefaultNodeInfo{
		ProtocolVersion: defaultProtocolVersion,
		DefaultNodeID:   rp.Addr().ID,
		ListenAddr:      rp.listener.Addr().String(),
		Network:         "testing",
		Version:         "1.2.3-rc0-deadbeef",
		Channels:        rp.channels,
		Moniker:         "remote_peer",
	}
}

func TestGetChIdStr(t *testing.T) {
	var i byte
	for ; i < math.MaxUint8; i++ {
		require.EqualValues(t, fmt.Sprintf("%#x", i), getChIdStr(i))
	}
}

const hextable = "0123456789abcdef"

func chidToHex(v byte) string {
	size := 3
	if v > 15 {
		size += 1
	}
	ret := make([]byte, size)
	ret[0] = '0'
	ret[1] = 'x'
	if v > 15 {
		ret[2] = hextable[v>>4]
		ret[3] = hextable[v&0x0f]
	} else {
		ret[2] = hextable[v]
	}
	return amino.BytesToStr(ret)
}

func BenchmarkChIdFormat(b *testing.B) {
	b.Run("fmt", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%#x", byte(i))
		}
	})
	b.Run("hex", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = chidToHex(byte(i))
		}
	})
	b.Run("table", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = getChIdStr(byte(i))
		}
	})
}
