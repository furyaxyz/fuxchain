package keeper_test

import (
	sdk "github.com/exfury/fuxchain/libs/cosmos-sdk/types"
	"github.com/exfury/fuxchain/libs/ibc-go/modules/apps/27-interchain-accounts/host/types"
)

func (suite *KeeperTestSuite) TestQueryParams() {
	ctx := sdk.WrapSDKContext(suite.chainA.GetContext())
	expParams := types.DefaultParams()
	res, _ := suite.chainA.GetSimApp().ICAHostKeeper.Params(ctx, &types.QueryParamsRequest{})
	suite.Require().Equal(&expParams, res.Params)
}
