package client

import (
	cliContext "github.com/exfury/fuxchain/libs/cosmos-sdk/client/context"
	"github.com/exfury/fuxchain/libs/cosmos-sdk/types/rest"
	"github.com/exfury/fuxchain/libs/ibc-go/modules/core/02-client/client/cli"
	govclient "github.com/exfury/fuxchain/x/gov/client"
	govrest "github.com/exfury/fuxchain/x/gov/client/rest"
	"net/http"
)

var (
	UpdateClientProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpdateClientProposal, emptyRestHandler)
)

func emptyRestHandler(ctx cliContext.CLIContext) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "unsupported-ibc-client",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Legacy REST Routes are not supported for IBC proposals")
		},
	}
}
