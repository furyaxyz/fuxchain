package types

import (
	codectypes "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/codec/types"
	"github.com/furyaxyz/fuxchain/libs/ibc-go/modules/core/exported"
)

// RegisterInterfaces register the ibc interfaces submodule implementations to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*exported.ClientState)(nil),
		&ClientState{},
	)
}
