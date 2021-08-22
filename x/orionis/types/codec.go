package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"

	// sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	exportedOracle "github.com/relevant-community/oracle/x/oracle/exported"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgOnrInr{}, "orionis/OnrInr", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	// registry.RegisterImplementations((*sdk.Msg)(nil),
	// 	&MsgOnrInr{},
	// )
	registry.RegisterInterface(
		"relevantcommunity.oracle.oracle.Claim",
		(*exportedOracle.Claim)(nil),
		&MsgOnrInr{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
