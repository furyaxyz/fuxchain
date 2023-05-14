package keeper

import (
	sdk "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/types"
	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/x/supply/exported"

	"github.com/furyaxyz/fuxchain/x/distribution/types"
)

// GetDistributionAccount returns the distribution ModuleAccount
func (k Keeper) GetDistributionAccount(ctx sdk.Context) exported.ModuleAccountI {
	return k.supplyKeeper.GetModuleAccount(ctx, types.ModuleName)
}
