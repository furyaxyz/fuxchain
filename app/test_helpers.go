package app

import (
	"github.com/spf13/viper"

	"github.com/exfury/fuxchain/libs/cosmos-sdk/codec"
	sdk "github.com/exfury/fuxchain/libs/cosmos-sdk/types"
	abci "github.com/exfury/fuxchain/libs/tendermint/abci/types"
	"github.com/exfury/fuxchain/libs/tendermint/libs/log"
	"github.com/exfury/fuxchain/libs/tendermint/types"
	dbm "github.com/exfury/fuxchain/libs/tm-db"
)

type Option func(option *SetupOption)

type SetupOption struct {
	chainId string
}

func WithChainId(chainId string) Option {
	return func(option *SetupOption) {
		option.chainId = chainId
	}
}

// Setup initializes a new FURYChainApp. A Nop logger is set in FURYChainApp.
func Setup(isCheckTx bool, options ...Option) *FURYChainApp {
	viper.Set(sdk.FlagDBBackend, string(dbm.MemDBBackend))
	types.DBBackend = string(dbm.MemDBBackend)
	db := dbm.NewMemDB()
	app := NewFURYChainApp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, 0)

	if !isCheckTx {
		setupOption := &SetupOption{chainId: ""}
		for _, opt := range options {
			opt(setupOption)
		}
		// init chain must be called to stop deliverState from being nil
		genesisState := NewDefaultGenesisState()
		stateBytes, err := codec.MarshalJSONIndent(app.Codec(), genesisState)
		if err != nil {
			panic(err)
		}

		// Initialize the chain
		app.InitChain(
			abci.RequestInitChain{
				Validators:    []abci.ValidatorUpdate{},
				AppStateBytes: stateBytes,
				ChainId:       setupOption.chainId,
			},
		)
	}

	return app
}
