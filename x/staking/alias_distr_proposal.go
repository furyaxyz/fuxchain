// nolint
// ALIASGEN: github.com/exfury/fuxchain/x/staking/types
package staking

import (
	"github.com/exfury/fuxchain/x/staking/types"
)

var (
	// functions aliases
	NewCommissionRates                = types.NewCommissionRates
	NewMsgEditValidatorCommissionRate = types.NewMsgEditValidatorCommissionRate
	NewMsgDestroyValidator            = types.NewMsgDestroyValidator
	NewMsgRegProxy                    = types.NewMsgRegProxy
	NewMsgBindProxy                   = types.NewMsgBindProxy
	NewMsgUnbindProxy                 = types.NewMsgUnbindProxy
)
