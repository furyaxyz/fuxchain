package types

import (
	sdk "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/types"
	sdkerrors "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/types/errors"
	"github.com/furyaxyz/fuxchain/libs/system"
)

// for denom convert wei to fury and reject fury direct
func (m *MsgTransfer) RulesFilter() (sdk.Msg, error) {
	if m.Token.Denom == sdk.DefaultBondDenom {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "ibc MsgTransfer not support "+system.Currency+" denom")
	}
	return m, nil
}
