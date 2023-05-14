package types

import (
	sdkerrors "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/types/errors"
)

var (
	ErrIBCAccountAlreadyExist = sdkerrors.Register(ModuleName, 2, "interchain account already registered")
	ErrIBCAccountNotExist     = sdkerrors.Register(ModuleName, 3, "interchain account not exist")
)
