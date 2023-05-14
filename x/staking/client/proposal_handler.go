package client

import (
	"github.com/exfury/fuxchain/x/staking/client/cli"
	"github.com/exfury/fuxchain/x/staking/client/rest"
	govcli "github.com/exfury/fuxchain/x/gov/client"
)

var (
	ProposeValidatorProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdProposeValidatorProposal,
		rest.ProposeValidatorProposalRESTHandler,
	)
)

