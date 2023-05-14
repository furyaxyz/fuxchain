// nolint:deadcode,unused
package params

import (
	abci "github.com/exfury/fuxchain/libs/tendermint/abci/types"
	"github.com/exfury/fuxchain/libs/tendermint/libs/log"
	dbm "github.com/exfury/fuxchain/libs/tm-db"

	"github.com/exfury/fuxchain/libs/cosmos-sdk/codec"
	"github.com/exfury/fuxchain/libs/cosmos-sdk/store"
	sdk "github.com/exfury/fuxchain/libs/cosmos-sdk/types"
)

type invalid struct{}

type s struct {
	I int
}

func createTestCodec() *codec.Codec {
	cdc := codec.New()
	sdk.RegisterCodec(cdc)
	cdc.RegisterConcrete(s{}, "test/s", nil)
	cdc.RegisterConcrete(invalid{}, "test/invalid", nil)
	return cdc
}

func defaultContext(key sdk.StoreKey, tkey sdk.StoreKey) sdk.Context {
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	cms.MountStoreWithDB(tkey, sdk.StoreTypeTransient, db)
	err := cms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}
	ctx := sdk.NewContext(cms, abci.Header{}, false, log.NewNopLogger())
	return ctx
}

func testComponents() (*codec.Codec, sdk.Context, sdk.StoreKey, sdk.StoreKey, Keeper) {
	cdc := createTestCodec()
	mkey := sdk.NewKVStoreKey("test")
	tkey := sdk.NewTransientStoreKey("transient_test")
	ctx := defaultContext(mkey, tkey)
	keeper := NewKeeper(cdc, mkey, tkey)

	return cdc, ctx, mkey, tkey, keeper
}
