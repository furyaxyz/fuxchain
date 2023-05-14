package transfer

import (
	"github.com/furyaxyz/fuxchain/libs/ibc-go/modules/apps/transfer/keeper"
	"github.com/furyaxyz/fuxchain/libs/ibc-go/modules/apps/transfer/types"
)

var (
	NewKeeper  = keeper.NewKeeper
	ModuleCdc  = types.ModuleCdc
	SetMarshal = types.SetMarshal
)
