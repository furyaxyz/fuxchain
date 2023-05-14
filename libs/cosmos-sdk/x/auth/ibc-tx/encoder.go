package ibc_tx

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/codec"
	codectypes "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/codec/types"
	sdk "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/types"
	sdkerrors "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/types/errors"
	ibctx "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/types/ibc-adapter"
	"github.com/furyaxyz/fuxchain/libs/cosmos-sdk/x/auth/types"
)

const MaxGasWanted = uint64((1 << 63) - 1)

func DefaultTxEncoder() ibctx.TxEncoder {
	return func(tx ibctx.Tx) ([]byte, error) {
		txWrapper, ok := tx.(*wrapper)
		if !ok {
			return nil, fmt.Errorf("expected %T, got %T", &wrapper{}, tx)
		}

		raw := &types.TxRaw{
			BodyBytes:     txWrapper.getBodyBytes(),
			AuthInfoBytes: txWrapper.getAuthInfoBytes(),
			Signatures:    txWrapper.tx.Signatures,
		}

		return proto.Marshal(raw)
	}
}

// DefaultJSONTxEncoder returns a default protobuf JSON TxEncoder using the provided Marshaler.
func DefaultJSONTxEncoder(cdc codec.ProtoCodecMarshaler) ibctx.IBCTxEncoder {
	return func(tx ibctx.Tx) ([]byte, error) {
		txWrapper, ok := tx.(*wrapper)
		if ok {
			return cdc.MarshalJSON(txWrapper.tx)
		}

		protoTx, ok := tx.(*TxAdapter)
		if ok {
			return cdc.MarshalJSON(protoTx)
		}

		return nil, fmt.Errorf("expected %T, got %T", &wrapper{}, tx)

	}
}

type TxAdapter struct {
	*types.Tx
}

func (t *TxAdapter) GetMsgs() []ibctx.Msg {
	if t == nil || t.Body == nil {
		return nil
	}

	anys := t.Body.Messages
	res := make([]ibctx.Msg, len(anys))
	for i, any := range anys {
		cached := any.GetCachedValue()
		if cached == nil {
			panic("Any cached value is nil. Transaction messages must be correctly packed Any values.")
		}
		res[i] = cached.(ibctx.Msg)
	}
	return res
}

// ValidateBasic implements the ValidateBasic method on sdk.Tx.
func (t *TxAdapter) ValidateBasic() error {
	if t == nil {
		return fmt.Errorf("bad Tx")
	}

	body := t.Body
	if body == nil {
		return fmt.Errorf("missing TxBody")
	}

	authInfo := t.AuthInfo
	if authInfo == nil {
		return fmt.Errorf("missing AuthInfo")
	}

	fee := authInfo.Fee
	if fee == nil {
		return fmt.Errorf("missing fee")
	}

	if fee.GasLimit > MaxGasWanted {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"invalid gas supplied; %d > %d", fee.GasLimit, MaxGasWanted,
		)
	}

	if fee.Amount.IsAnyNil() {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFee,
			"invalid fee provided: null",
		)
	}

	if fee.Amount.IsAnyNegative() {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFee,
			"invalid fee provided: %s", fee.Amount,
		)
	}

	if fee.Payer != "" {
		_, err := sdk.AccAddressFromBech32(fee.Payer)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid fee payer address (%s)", err)
		}
	}

	sigs := t.Signatures

	if len(sigs) == 0 {
		return sdkerrors.ErrNoSignatures
	}

	if len(sigs) != len(t.GetSigners()) {
		return sdkerrors.Wrapf(
			sdkerrors.ErrUnauthorized,
			"wrong number of signers; expected %d, got %d", len(t.GetSigners()), len(sigs),
		)
	}

	return nil
}

// GetSigners retrieves all the signers of a tx.
// This includes all unique signers of the messages (in order),
// as well as the FeePayer (if specified and not already included).
func (t *TxAdapter) GetSigners() []sdk.AccAddress {
	var signers []sdk.AccAddress
	seen := map[string]bool{}

	for _, msg := range t.GetMsgs() {
		for _, addr := range msg.GetSigners() {
			if !seen[addr.String()] {
				signers = append(signers, addr)
				seen[addr.String()] = true
			}
		}
	}

	// ensure any specified fee payer is included in the required signers (at the end)
	feePayer := t.AuthInfo.Fee.Payer
	if feePayer != "" && !seen[feePayer] {
		payerAddr, err := sdk.AccAddressFromBech32(feePayer)
		if err != nil {
			panic(err)
		}
		signers = append(signers, payerAddr)
		seen[feePayer] = true
	}

	return signers
}

//
//func (t *TxRawAdapter) GetGas() uint64 {
//	return t.AuthInfo.Fee.GasLimit
//}
//func (t *TxRawAdapter) GetFee() sdk.CoinAdapters {
//	return t.AuthInfo.Fee.Amount
//}
//func (t *TxRawAdapter) FeePayer() sdk.AccAddress {
//	feePayer := t.AuthInfo.Fee.Payer
//	if feePayer != "" {
//		payerAddr, err := sdk.AccAddressFromBech32(feePayer)
//		if err != nil {
//			panic(err)
//		}
//		return payerAddr
//	}
//	// use first signer as default if no payer specified
//	return t.GetSigners()[0]
//}
//
//func (t *TxRawAdapter) FeeGranter() sdk.AccAddress {
//	feePayer := t.AuthInfo.Fee.Granter
//	if feePayer != "" {
//		granterAddr, err := sdk.AccAddressFromBech32(feePayer)
//		if err != nil {
//			panic(err)
//		}
//		return granterAddr
//	}
//	return nil
//}
//
//// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
//func (t *TxRawAdapter) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
//	if t.Body != nil {
//		if err := t.Body.UnpackInterfaces(unpacker); err != nil {
//			return err
//		}
//	}
//
//	if t.AuthInfo != nil {
//		return t.AuthInfo.UnpackInterfaces(unpacker)
//	}
//
//	return nil
//}
//
//// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
//func (m *TxRawAdapter) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
//	for _, any := range m.Messages {
//		var msg sdk.Msg
//		err := unpacker.UnpackAny(any, &msg)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
//func (m *TxRawAdapter) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
//	for _, signerInfo := range m.SignerInfos {
//		err := signerInfo.UnpackInterfaces(unpacker)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
//func (m *TxRawAdapter) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
//	return unpacker.UnpackAny(m.PublicKey, new(cryptotypes.PubKey))
//}

//TODO add call RegisterInterfaces
// RegisterInterfaces registers the sdk.Tx interface.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface("cosmos.tx.v1beta1.Tx", (*sdk.Tx)(nil))
	registry.RegisterImplementations((*sdk.Tx)(nil), &types.Tx{})
}
