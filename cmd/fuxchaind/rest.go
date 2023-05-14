package main

import (
	"github.com/furyaxyz/fuxchain/app"
	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/types/tx"
	"github.com/furyaxyz/fuxchain/x/wasm/proxy"

	mintclient "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/x/mint/client"
	mintrest "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/x/mint/client/rest"
	erc20client "github.com/furyaxyz/fuxchain/x/erc20/client"
	erc20rest "github.com/furyaxyz/fuxchain/x/erc20/client/rest"
	evmclient "github.com/furyaxyz/fuxchain/x/evm/client"

	"github.com/furyaxyz/fuxchain/app/rpc"
	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/client"
	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/client/lcd"
	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/x/auth"
	authrest "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/x/auth/client/rest"
	bankrest "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/x/bank/client/rest"
	supplyrest "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/x/supply/client/rest"
	dist "github.com/furyaxyz/fuxchain/x/distribution"
	distr "github.com/furyaxyz/fuxchain/x/distribution"
	distrest "github.com/furyaxyz/fuxchain/x/distribution/client/rest"
	evmrest "github.com/furyaxyz/fuxchain/x/evm/client/rest"
	fsrest "github.com/furyaxyz/fuxchain/x/feesplit/client/rest"
	govrest "github.com/furyaxyz/fuxchain/x/gov/client/rest"
	paramsclient "github.com/furyaxyz/fuxchain/x/params/client"
	stakingclient "github.com/furyaxyz/fuxchain/x/staking/client"
	stakingrest "github.com/furyaxyz/fuxchain/x/staking/client/rest"
	"github.com/furyaxyz/fuxchain/x/token"
	tokensrest "github.com/furyaxyz/fuxchain/x/token/client/rest"
	wasmrest "github.com/furyaxyz/fuxchain/x/wasm/client/rest"
)

// registerRoutes registers the routes from the different modules for the LCD.
// NOTE: details on the routes added for each module are in the module documentation
// NOTE: If making updates here you also need to update the test helper in client/lcd/test_helper.go
func registerRoutes(rs *lcd.RestServer) {
	registerGrpc(rs)
	rpc.RegisterRoutes(rs)
	registerRoutesV1(rs)
	registerRoutesV2(rs)
	proxy.SetCliContext(rs.CliCtx)
}

func registerGrpc(rs *lcd.RestServer) {
	app.ModuleBasics.RegisterGRPCGatewayRoutes(rs.CliCtx, rs.GRPCGatewayRouter)
	app.ModuleBasics.RegisterRPCRouterForGRPC(rs.CliCtx, rs.Mux)
	tx.RegisterGRPCGatewayRoutes(rs.CliCtx, rs.GRPCGatewayRouter)
}

func registerRoutesV1(rs *lcd.RestServer) {
	v1Router := rs.Mux.PathPrefix("/v1").Name("v1").Subrouter()
	client.RegisterRoutes(rs.CliCtx, v1Router)
	authrest.RegisterRoutes(rs.CliCtx, v1Router, auth.StoreKey)
	bankrest.RegisterRoutes(rs.CliCtx, v1Router)
	stakingrest.RegisterRoutes(rs.CliCtx, v1Router)
	distrest.RegisterRoutes(rs.CliCtx, v1Router, dist.StoreKey)

	tokensrest.RegisterRoutes(rs.CliCtx, v1Router, token.StoreKey)
	supplyrest.RegisterRoutes(rs.CliCtx, v1Router)
	evmrest.RegisterRoutes(rs.CliCtx, v1Router)
	erc20rest.RegisterRoutes(rs.CliCtx, v1Router)
	wasmrest.RegisterRoutes(rs.CliCtx, v1Router)
	fsrest.RegisterRoutes(rs.CliCtx, v1Router)
	govrest.RegisterRoutes(rs.CliCtx, v1Router,
		[]govrest.ProposalRESTHandler{
			paramsclient.ProposalHandler.RESTHandler(rs.CliCtx),
			distr.CommunityPoolSpendProposalHandler.RESTHandler(rs.CliCtx),
			distr.ChangeDistributionTypeProposalHandler.RESTHandler(rs.CliCtx),
			distr.WithdrawRewardEnabledProposalHandler.RESTHandler(rs.CliCtx),
			distr.RewardTruncatePrecisionProposalHandler.RESTHandler(rs.CliCtx),
			evmclient.ManageContractDeploymentWhitelistProposalHandler.RESTHandler(rs.CliCtx),
			evmclient.ManageSysContractAddressProposalHandler.RESTHandler(rs.CliCtx),
			evmclient.ManageContractByteCodeProposalHandler.RESTHandler(rs.CliCtx),
			mintclient.ManageTreasuresProposalHandler.RESTHandler(rs.CliCtx),
			mintclient.ExtraProposalHandler.RESTHandler(rs.CliCtx),
			erc20client.TokenMappingProposalHandler.RESTHandler(rs.CliCtx),
			stakingclient.ProposeValidatorProposalHandler.RESTHandler(rs.CliCtx),
		},
	)
	mintrest.RegisterRoutes(rs.CliCtx, v1Router)

}

func registerRoutesV2(rs *lcd.RestServer) {
	v2Router := rs.Mux.PathPrefix("/v2").Name("v2").Subrouter()
	client.RegisterRoutes(rs.CliCtx, v2Router)
	authrest.RegisterRoutes(rs.CliCtx, v2Router, auth.StoreKey)
	bankrest.RegisterRoutes(rs.CliCtx, v2Router)
	stakingrest.RegisterRoutes(rs.CliCtx, v2Router)
	distrest.RegisterRoutes(rs.CliCtx, v2Router, dist.StoreKey)
	tokensrest.RegisterRoutesV2(rs.CliCtx, v2Router, token.StoreKey)
	fsrest.RegisterRoutesV2(rs.CliCtx, v2Router)
}
