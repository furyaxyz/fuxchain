package keeper

import (
	sdk "github.com/exfury/fuxchain/libs/cosmos-sdk/types"
	govtypes "github.com/exfury/fuxchain/x/gov/types"
)

// GovKeeper defines the expected gov Keeper
type GovKeeper interface {
	GetDepositParams(ctx sdk.Context) govtypes.DepositParams
	GetVotingParams(ctx sdk.Context) govtypes.VotingParams
}
