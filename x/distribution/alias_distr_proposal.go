// nolint
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/furyaxyz/fuxchain/x/distribution/types
// ALIASGEN: github.com/furyaxyz/fuxchain/x/distribution/client
package distribution

import (
	"github.com/furyaxyz/fuxchain/x/distribution/client"
	"github.com/furyaxyz/fuxchain/x/distribution/types"
)

var (
	NewMsgWithdrawDelegatorReward          = types.NewMsgWithdrawDelegatorReward
	CommunityPoolSpendProposalHandler      = client.CommunityPoolSpendProposalHandler
	ChangeDistributionTypeProposalHandler  = client.ChangeDistributionTypeProposalHandler
	WithdrawRewardEnabledProposalHandler   = client.WithdrawRewardEnabledProposalHandler
	RewardTruncatePrecisionProposalHandler = client.RewardTruncatePrecisionProposalHandler
	NewMsgWithdrawDelegatorAllRewards      = types.NewMsgWithdrawDelegatorAllRewards
)
