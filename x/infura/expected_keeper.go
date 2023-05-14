package infura

import evm "github.com/furyaxyz/fuxchain/x/evm/watcher"

type EvmKeeper interface {
	SetObserverKeeper(keeper evm.InfuraKeeper)
}
