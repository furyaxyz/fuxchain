package types

import (
	sdk "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/types"
	authexported "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/x/auth/exported"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authexported.Account
	IterateAccounts(ctx sdk.Context, cb func(account authexported.Account) bool)
}
