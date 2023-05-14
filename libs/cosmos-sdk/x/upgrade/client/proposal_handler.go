package client

import (
	govclient "github.com/exfury/fuxchain/libs/cosmos-sdk/x/gov/client"
	"github.com/exfury/fuxchain/libs/cosmos-sdk/x/upgrade/client/cli"
	"github.com/exfury/fuxchain/libs/cosmos-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
