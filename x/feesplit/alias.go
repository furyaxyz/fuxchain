package feesplit

import (
	"github.com/exfury/fuxchain/x/feesplit/keeper"
	"github.com/exfury/fuxchain/x/feesplit/types"
)

const (
	ModuleName = types.ModuleName
	StoreKey   = types.StoreKey
	RouterKey  = types.RouterKey
)

var (
	NewKeeper           = keeper.NewKeeper
	SetParamsNeedUpdate = types.SetParamsNeedUpdate
)

type (
	Keeper = keeper.Keeper
)
