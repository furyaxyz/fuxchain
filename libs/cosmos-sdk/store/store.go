package store

import (
	dbm "github.com/furyaxyz/fuxchain/libs/tm-db"

	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/store/cache"
	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/store/rootmulti"
	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
