package client

import (
	"github.com/furyaxyz/fuxchain/x/staking/client/cli"
	"github.com/furyaxyz/fuxchain/x/staking/client/rest"
	govcli "github.com/furyaxyz/fuxchain/x/gov/client"
)

var (
	ProposeValidatorProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdProposeValidatorProposal,
		rest.ProposeValidatorProposalRESTHandler,
	)
)

