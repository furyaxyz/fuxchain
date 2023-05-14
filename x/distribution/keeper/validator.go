package keeper

import (
	sdk "github.com/exfury/fuxchain/libs/cosmos-sdk/types"

	"github.com/exfury/fuxchain/x/staking/exported"
)

// initialize rewards for a new validator
func (k Keeper) initializeValidator(ctx sdk.Context, val exported.ValidatorI) {
	k.initializeValidatorDistrProposal(ctx, val)
	return
}
