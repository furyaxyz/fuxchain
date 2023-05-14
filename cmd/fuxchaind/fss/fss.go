package fss

import (
	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/server"
	sdk "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/types"
	"github.com/furyaxyz/fuxchain/libs/iavl"
	tmtypes "github.com/furyaxyz/fuxchain/libs/tendermint/types"
	"github.com/spf13/cobra"
)

const (
	flagDataDir = "data_dir"
)

func Command(ctx *server.Context) *cobra.Command {
	iavl.SetLogger(ctx.Logger.With("module", "iavl"))
	return fssCmd
}

var fssCmd = &cobra.Command{
	Use:   "fss",
	Short: "FSS is an auxiliary fast storage system to Tree",
	Long: `tree fast storage related commands:
This command include a set of command of the Tree fast storage.
include create sub command`,
}

func init() {
	fssCmd.PersistentFlags().StringP(flagDataDir, "d", "./", "The chain data file location")
	fssCmd.PersistentFlags().String(sdk.FlagDBBackend, tmtypes.DBBackend, "Database backend: goleveldb | rocksdb")
}
