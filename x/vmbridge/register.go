package vmbridge

import (
	"github.com/exfury/fuxchain/libs/cosmos-sdk/codec"
	"github.com/exfury/fuxchain/libs/cosmos-sdk/types/module"
	"github.com/exfury/fuxchain/x/vmbridge/keeper"
	"github.com/exfury/fuxchain/x/wasm"
)

func RegisterServices(cfg module.Configurator, keeper keeper.Keeper) {
	RegisterMsgServer(cfg.MsgServer(), NewMsgServerImpl(keeper))
}

func GetWasmOpts(cdc *codec.ProtoCodec) wasm.Option {
	return wasm.WithMessageEncoders(RegisterSendToEvmEncoder(cdc))
}
