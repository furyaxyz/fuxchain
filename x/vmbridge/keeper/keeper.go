package keeper

import (
	"github.com/exfury/fuxchain/x/vmbridge/types"

	"github.com/exfury/fuxchain/libs/cosmos-sdk/codec"
	"github.com/exfury/fuxchain/libs/tendermint/libs/log"
)

type Keeper struct {
	cdc *codec.CodecProxy

	logger log.Logger

	evmKeeper     EVMKeeper
	wasmKeeper    WASMKeeper
	accountKeeper AccountKeeper
}

func NewKeeper(cdc *codec.CodecProxy, logger log.Logger, evmKeeper EVMKeeper, wasmKeeper WASMKeeper, accountKeeper AccountKeeper) *Keeper {
	logger = logger.With("module", types.ModuleName)
	return &Keeper{cdc: cdc, logger: logger, evmKeeper: evmKeeper, wasmKeeper: wasmKeeper, accountKeeper: accountKeeper}
}

func (k Keeper) Logger() log.Logger {
	return k.logger
}

func (k Keeper) getAminoCodec() *codec.Codec {
	return k.cdc.GetCdc()
}

func (k Keeper) GetProtoCodec() *codec.ProtoCodec {
	return k.cdc.GetProtocMarshal()
}
