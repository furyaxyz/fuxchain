package consensus

import (
	amino "github.com/tendermint/go-amino"

	"github.com/exfury/fuxchain/libs/tendermint/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterMessages(cdc)
	RegisterWALMessages(cdc)
	types.RegisterBlockAmino(cdc)
}
