package client

import (
	govclient "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/x/gov/client"
	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/x/upgrade/client/cli"
	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
