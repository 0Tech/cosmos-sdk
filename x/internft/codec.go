package internft

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
	govcodec "github.com/cosmos/cosmos-sdk/x/gov/codec"
)

// RegisterLegacyAminoCodec registers concrete types on the LegacyAmino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	for msg, name := range map[sdk.Msg]string{
		&MsgSend{}:        "MsgSend",
		&MsgNewClass{}:    "MsgNewClass",
		&MsgUpdateClass{}: "MsgUpdateClass",
		&MsgMintNFT{}:     "MsgMintNFT",
		&MsgBurnNFT{}:     "MsgBurnNFT",
		&MsgUpdateNFT{}:   "MsgUpdateNFT",
	} {
		const prefix = "cosmos-sdk/x/internft/"
		legacy.RegisterAminoMsg(cdc, msg, prefix+name)
	}
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSend{},
		&MsgNewClass{},
		&MsgUpdateClass{},
		&MsgMintNFT{},
		&MsgBurnNFT{},
		&MsgUpdateNFT{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func init() {
	// Register all Amino interfaces and concrete types on the authz and gov Amino codec so that this can later be
	// used to properly serialize MsgGrant, MsgExec and MsgSubmitProposal instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
	RegisterLegacyAminoCodec(govcodec.Amino)
}
