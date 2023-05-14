package mpt

import (
	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/server"
	sdk "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/types"
	tmtypes "github.com/furyaxyz/fuxchain/libs/tendermint/types"
	"github.com/spf13/cobra"
)

func MptCmd(ctx *server.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mpt",
		Short: "migrate iavl state to mpt state (if use migrate mpt data, then you should set `--use-composite-key true` when you decide to use mpt to store the coming data)",
	}

	cmd.AddCommand(
		iavl2mptCmd(ctx),
		cleanIavlStoreCmd(ctx),
		mptViewerCmd(ctx),
		AccountGetCmd(ctx),
		genSnapCmd(ctx),
	)

	cmd.PersistentFlags().String(sdk.FlagDBBackend, tmtypes.DBBackend, "Database backend: goleveldb | rocksdb")

	return cmd
}
