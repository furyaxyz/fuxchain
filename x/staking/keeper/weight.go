package keeper

import (
	sdk "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/types"
	"github.com/furyaxyz/fuxchain/x/staking/types"
)

func calculateWeight(tokens sdk.Dec) types.Shares {
	return tokens
}

func SimulateWeight(tokens sdk.Dec) types.Shares {
	return calculateWeight(tokens)
}
