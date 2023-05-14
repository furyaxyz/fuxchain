package types

import (
	interfacetypes "github.com/exfury/fuxchain/libs/cosmos-sdk/codec/types"
	txmsg "github.com/exfury/fuxchain/libs/cosmos-sdk/types/ibc-adapter"
	"github.com/exfury/fuxchain/libs/cosmos-sdk/types/msgservice"
)

func RegisterInterface(registry interfacetypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*txmsg.Msg)(nil),
		&MsgSendToEvm{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
