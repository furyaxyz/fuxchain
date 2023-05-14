package client

import (
	"github.com/exfury/fuxchain/x/feesplit/client/cli"
	"github.com/exfury/fuxchain/x/feesplit/client/rest"
	govcli "github.com/exfury/fuxchain/x/gov/client"
)

var (
	// FeeSplitSharesProposalHandler alias gov NewProposalHandler
	FeeSplitSharesProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdFeeSplitSharesProposal,
		rest.FeeSplitSharesProposalRESTHandler,
	)
)
