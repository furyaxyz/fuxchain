package types

import (
	sdk "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/types"
)

// GetDelegatorAddress gets delegator address
func (d Delegator) GetDelegatorAddress() sdk.AccAddress {
	return d.DelegatorAddress
}
