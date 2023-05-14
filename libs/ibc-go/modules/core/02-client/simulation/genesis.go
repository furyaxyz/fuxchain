package simulation

import (
	"math/rand"

	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/x/simulation"
	"github.com/furyaxyz/fuxchain/libs/ibc-go/modules/core/02-client/types"
)

// GenClientGenesis returns the default client genesis state.
func GenClientGenesis(_ *rand.Rand, _ []simulation.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
