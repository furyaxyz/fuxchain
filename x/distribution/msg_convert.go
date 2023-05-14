package distribution

import (
	"encoding/json"
	"errors"
	"github.com/exfury/fuxchain/libs/system"

	"github.com/exfury/fuxchain/libs/cosmos-sdk/baseapp"
	sdk "github.com/exfury/fuxchain/libs/cosmos-sdk/types"
	"github.com/exfury/fuxchain/x/common"
	"github.com/exfury/fuxchain/x/distribution/types"
)

var (
	ErrCheckSignerFail = errors.New("check signer fail")
)

func init() {
	RegisterConvert()
}

func RegisterConvert() {
	baseapp.RegisterCmHandle(system.Chain+"/distribution/MsgWithdrawDelegatorAllRewards", baseapp.NewCMHandle(ConvertWithdrawDelegatorAllRewardsMsg, 0))
}

func ConvertWithdrawDelegatorAllRewardsMsg(data []byte, signers []sdk.AccAddress) (sdk.Msg, error) {
	newMsg := types.MsgWithdrawDelegatorAllRewards{}
	err := json.Unmarshal(data, &newMsg)
	if err != nil {
		return nil, err
	}
	err = newMsg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	if ok := common.CheckSignerAddress(signers, newMsg.GetSigners()); !ok {
		return nil, ErrCheckSignerFail
	}
	return newMsg, nil
}
