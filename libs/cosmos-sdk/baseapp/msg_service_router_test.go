package baseapp_test

import (
	chaincodec "github.com/furyaxyz/fuxchain/app/codec"
	"github.com/furyaxyz/fuxchain/libs/ibc-go/testing/simapp"
	"github.com/furyaxyz/fuxchain/x/evm"
	"os"
	"testing"

	"github.com/furyaxyz/fuxchain/libs/tendermint/libs/log"

	dbm "github.com/furyaxyz/fuxchain/libs/tm-db"
	"github.com/stretchr/testify/require"

	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/baseapp"

	"github.com/furyaxyz/fuxchain/x/evm/types/testdata"
)

func TestRegisterMsgService(t *testing.T) {
	db := dbm.NewMemDB()

	// Create an encoding config that doesn't register testdata Msg services.
	codecProxy, interfaceRegistry := chaincodec.MakeCodecSuit(simapp.ModuleBasics)
	app := baseapp.NewBaseApp("test", log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, evm.TxDecoder(codecProxy))
	app.SetInterfaceRegistry(interfaceRegistry)
	require.Panics(t, func() {
		testdata.RegisterMsgServer(
			app.MsgServiceRouter(),
			testdata.MsgServerImpl{},
		)
	})

	// Register testdata Msg services, and rerun `RegisterService`.
	testdata.RegisterInterfaces(interfaceRegistry)
	require.NotPanics(t, func() {
		testdata.RegisterMsgServer(
			app.MsgServiceRouter(),
			testdata.MsgServerImpl{},
		)
	})
}

func TestRegisterMsgServiceTwice(t *testing.T) {
	// Setup baseapp.
	db := dbm.NewMemDB()
	codecProxy, interfaceRegistry := chaincodec.MakeCodecSuit(simapp.ModuleBasics)
	app := baseapp.NewBaseApp("test", log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, evm.TxDecoder(codecProxy))
	app.SetInterfaceRegistry(interfaceRegistry)
	testdata.RegisterInterfaces(interfaceRegistry)

	// First time registering service shouldn't panic.
	require.NotPanics(t, func() {
		testdata.RegisterMsgServer(
			app.MsgServiceRouter(),
			testdata.MsgServerImpl{},
		)
	})

	// Second time should panic.
	require.Panics(t, func() {
		testdata.RegisterMsgServer(
			app.MsgServiceRouter(),
			testdata.MsgServerImpl{},
		)
	})
}
