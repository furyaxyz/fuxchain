package proto

import (
	"os"
	"testing"

	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/store"
	sdk "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/types"
	abci "github.com/furyaxyz/fuxchain/libs/tendermint/abci/types"
	"github.com/furyaxyz/fuxchain/libs/tendermint/libs/log"
	dbm "github.com/furyaxyz/fuxchain/libs/tm-db"
	"github.com/stretchr/testify/require"
)

func createTestInput(t *testing.T) (sdk.Context, ProtocolKeeper) {
	keyMain := sdk.NewKVStoreKey("main")

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyMain, sdk.StoreTypeIAVL, db)

	require.NoError(t, ms.LoadLatestVersion())

	ctx := sdk.NewContext(ms, abci.Header{}, false, log.NewTMLogger(os.Stdout))

	keeper := NewProtocolKeeper(keyMain)

	return ctx, keeper
}
