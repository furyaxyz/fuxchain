package client

import (
	"github.com/furyaxyz/fuxchain/x/feesplit/client/cli"
	"github.com/furyaxyz/fuxchain/x/feesplit/client/rest"
	govcli "github.com/furyaxyz/fuxchain/x/gov/client"
)

var (
	// FeeSplitSharesProposalHandler alias gov NewProposalHandler
	FeeSplitSharesProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdFeeSplitSharesProposal,
		rest.FeeSplitSharesProposalRESTHandler,
	)
)
