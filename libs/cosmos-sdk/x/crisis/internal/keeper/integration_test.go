package keeper_test

import (
	abci "github.com/furyaxyz/fuxchain/libs/tendermint/abci/types"
	"github.com/furyaxyz/fuxchain/libs/tendermint/libs/log"
	dbm "github.com/furyaxyz/fuxchain/libs/tm-db"

	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/codec"
	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/simapp"
)

func createTestApp() *simapp.SimApp {
	db := dbm.NewMemDB()
	app := simapp.NewSimApp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, 5)
	// init chain must be called to stop deliverState from being nil
	genesisState := simapp.NewDefaultGenesisState()
	stateBytes, err := codec.MarshalJSONIndent(app.Codec(), genesisState)
	if err != nil {
		panic(err)
	}

	// Initialize the chain
	app.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	app.Commit(abci.RequestCommit{})
	app.BeginBlock(abci.RequestBeginBlock{Header: abci.Header{Height: app.LastBlockHeight() + 1}})

	return app
}
