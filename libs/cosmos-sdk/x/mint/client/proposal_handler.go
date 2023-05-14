package client

import (
	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/x/mint/client/cli"
	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/x/mint/client/rest"
	govcli "github.com/furyaxyz/fuxchain/x/gov/client"
)

var (
	ManageTreasuresProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdManageTreasuresProposal,
		rest.ManageTreasuresProposalRESTHandler,
	)

	ExtraProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdExtraProposal,
		rest.ExtraProposalRESTHandler,
	)
)
