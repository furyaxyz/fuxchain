BIN_NAME=fuxchaind
FUCHAIN_TOP=${GOPATH}/src/github.com/furyaxyz/fuxchain
FUCHAIN_BIN=${FUCHAIN_TOP}/build
FUCHAIN_BIN=${GOPATH}/bin
FUCHAIN_NET_TOP=`pwd`
FUCHAIN_NET_CACHE=${FUCHAIN_NET_TOP}/cache
CHAIN_ID="clockend-4200"


BASE_PORT_PREFIX=26600
P2P_PORT_SUFFIX=56
RPC_PORT_SUFFIX=57
REST_PORT=8545
let BASE_PORT=${BASE_PORT_PREFIX}+${P2P_PORT_SUFFIX}
let seedp2pport=${BASE_PORT_PREFIX}+${P2P_PORT_SUFFIX}
let seedrpcport=${BASE_PORT_PREFIX}+${RPC_PORT_SUFFIX}
let seedrestport=${seedrpcport}+1
