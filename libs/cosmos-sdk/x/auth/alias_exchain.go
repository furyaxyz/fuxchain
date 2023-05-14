package auth

import (
	"github.com/exfury/fuxchain/libs/cosmos-sdk/x/auth/exported"
	"github.com/exfury/fuxchain/libs/cosmos-sdk/x/auth/keeper"
)

type (
	Account       = exported.Account
	ModuleAccount = exported.ModuleAccount
	ObserverI     = keeper.ObserverI
)
