package vmbridge

import (
	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/codec"
	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/types/module"
	"github.com/furyaxyz/fuxchain/x/vmbridge/keeper"
	"github.com/furyaxyz/fuxchain/x/wasm"
)

func RegisterServices(cfg module.Configurator, keeper keeper.Keeper) {
	RegisterMsgServer(cfg.MsgServer(), NewMsgServerImpl(keeper))
}

func GetWasmOpts(cdc *codec.ProtoCodec) wasm.Option {
	return wasm.WithMessageEncoders(RegisterSendToEvmEncoder(cdc))
}
