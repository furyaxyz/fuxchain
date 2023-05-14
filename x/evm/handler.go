package evm

import (
	bam "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/baseapp"
	sdk "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/types"
	sdkerrors "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/types/errors"
	cfg "github.com/furyaxyz/fuxchain/libs/tendermint/config"
	"github.com/furyaxyz/fuxchain/x/evm/txs"
	"github.com/furyaxyz/fuxchain/x/evm/txs/base"
	"github.com/furyaxyz/fuxchain/x/evm/types"
	"github.com/furyaxyz/fuxchain/x/evm/watcher"
)

// NewHandler returns a handler for Ethermint type messages.
func NewHandler(k *Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (result *sdk.Result, err error) {
		ctx.SetEventManager(sdk.NewEventManager())

		if ctx.IsDeliver() {
			k.UpdatedAccount = k.UpdatedAccount[:0]
		}

		evmTx, ok := msg.(*types.MsgEthereumTx)
		if ok {
			if watcher.IsWatcherEnabled() {
				ctx.SetWatcher(watcher.NewTxWatcher())
			}
			result, err = handleMsgEthereumTx(ctx, k, evmTx)
			if err != nil {
				err = sdkerrors.New(types.ModuleName, types.CodeSpaceEvmCallFailed, err.Error())
			}
		} else {
			err = sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}

		return result, err
	}
}

func updateHGU(ctx sdk.Context, msg sdk.Msg) {
	if cfg.DynamicConfig.GetMaxGasUsedPerBlock() <= 0 {
		return
	}

	msgFnSignature, toDeployContractSize := getMsgCallFnSignature(msg)

	if msgFnSignature == nil {
		return
	}

	gc := int64(ctx.GasMeter().GasConsumed())
	if toDeployContractSize > 0 {
		// calculate average gas consume for deploy contract case
		gc = gc / int64(toDeployContractSize)
	}

	bam.InstanceOfHistoryGasUsedRecordDB().UpdateGasUsed(msgFnSignature, gc)
}

func getMsgCallFnSignature(msg sdk.Msg) ([]byte, int) {
	switch msg := msg.(type) {
	case *types.MsgEthereumTx:
		return msg.GetTxFnSignatureInfo()
	default:
		return nil, 0
	}
}

// handleMsgEthereumTx handles an Ethereum specific tx
// 1. txs can be divided into TraceTxLog/CheckTx/DeliverTx
func handleMsgEthereumTx(ctx sdk.Context, k *Keeper, msg *types.MsgEthereumTx) (*sdk.Result, error) {
	txFactory := txs.NewFactory(base.Config{
		Ctx:    ctx,
		Keeper: k,
	})
	tx, err := txFactory.CreateTx()
	if err != nil {
		return nil, err
	}
	defer tx.Dispose()

	// core logical to handle ethereum tx
	rst, err := txs.TransitionEvmTx(tx, msg)
	if err == nil && !ctx.IsCheckTx() {
		updateHGU(ctx, msg)
	}

	return rst, err
}
