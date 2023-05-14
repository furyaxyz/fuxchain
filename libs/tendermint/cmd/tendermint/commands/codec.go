package commands

import (
	amino "github.com/tendermint/go-amino"

	cryptoamino "github.com/furyaxyz/fuxchain/libs/tendermint/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoamino.RegisterAmino(cdc)
}
