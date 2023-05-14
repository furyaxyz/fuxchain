package types

import (
	sdkerrors "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/types/errors"
)

// x/crisis module sentinel errors
var (
	ErrNoSender         = sdkerrors.Register(ModuleName, 1, "sender address is empty")
	ErrUnknownInvariant = sdkerrors.Register(ModuleName, 2, "unknown invariant")
)
