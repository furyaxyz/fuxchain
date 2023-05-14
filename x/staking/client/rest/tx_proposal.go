package rest

import (
	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/client/context"
	govRest "github.com/furyaxyz/fuxchain/x/gov/client/rest"
)

// ProposeValidatorProposalRESTHandler defines propose validator proposal handler
func ProposeValidatorProposalRESTHandler(context.CLIContext) govRest.ProposalRESTHandler {
	return govRest.ProposalRESTHandler{}
}
