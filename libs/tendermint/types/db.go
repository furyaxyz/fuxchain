package types

import dbm "github.com/exfury/fuxchain/libs/tm-db"

// DBBackend This is set at compile time.
var DBBackend = string(dbm.GoLevelDBBackend)
