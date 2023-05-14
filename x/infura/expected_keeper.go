package infura

import evm "github.com/exfury/fuxchain/x/evm/watcher"

type EvmKeeper interface {
	SetObserverKeeper(keeper evm.InfuraKeeper)
}
