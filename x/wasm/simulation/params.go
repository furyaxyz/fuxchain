package simulation

import (
	"fmt"
	"math/rand"

	"github.com/exfury/fuxchain/libs/cosmos-sdk/codec"
	"github.com/exfury/fuxchain/libs/cosmos-sdk/x/simulation"
	simtypes "github.com/exfury/fuxchain/libs/cosmos-sdk/x/simulation"

	"github.com/exfury/fuxchain/x/wasm/types"
)

func ParamChanges(r *rand.Rand, cdc codec.Codec) []simtypes.ParamChange {
	params := types.DefaultParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.ParamStoreKeyUploadAccess),
			func(r *rand.Rand) string {
				jsonBz, err := cdc.MarshalJSON(&params.CodeUploadAccess)
				if err != nil {
					panic(err)
				}
				return string(jsonBz)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.ParamStoreKeyInstantiateAccess),
			func(r *rand.Rand) string {
				return fmt.Sprintf("%q", params.CodeUploadAccess.Permission.String())
			},
		),
	}
}
