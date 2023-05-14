package rest

import (
	"github.com/exfury/fuxchain/libs/cosmos-sdk/client/context"
	govRest "github.com/exfury/fuxchain/x/gov/client/rest"
)

// ProposeValidatorProposalRESTHandler defines propose validator proposal handler
func ProposeValidatorProposalRESTHandler(context.CLIContext) govRest.ProposalRESTHandler {
	return govRest.ProposalRESTHandler{}
}
