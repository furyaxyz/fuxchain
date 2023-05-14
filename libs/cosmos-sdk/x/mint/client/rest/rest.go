package rest

import (
	"github.com/gorilla/mux"
	govRest "github.com/exfury/fuxchain/x/gov/client/rest"

	"github.com/exfury/fuxchain/libs/cosmos-sdk/client/context"
)

// RegisterRoutes registers minting module REST handlers on the provided router.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
}

// ManageTreasuresProposalRESTHandler defines mint proposal handler
func ManageTreasuresProposalRESTHandler(context.CLIContext) govRest.ProposalRESTHandler {
	return govRest.ProposalRESTHandler{}
}

// ExtraProposalRESTHandler defines mint proposal handler
func ExtraProposalRESTHandler(context.CLIContext) govRest.ProposalRESTHandler {
	return govRest.ProposalRESTHandler{}
}
