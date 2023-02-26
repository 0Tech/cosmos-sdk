package internft_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
	"github.com/cosmos/cosmos-sdk/x/internft"
)

func TestMsgSend(t *testing.T) {
	addrs := createAddresses(2, "addr")
	classID := createClassIDs(1, "class")[0]

	testCases := map[string]struct {
		sender    sdk.AccAddress
		recipient sdk.AccAddress
		classID   string
		err       error
	}{
		"valid msg": {
			sender:    addrs[0],
			recipient: addrs[1],
			classID:   classID,
		},
		"invalid sender": {
			recipient: addrs[1],
			classID:   classID,
			err:       sdkerrors.ErrInvalidAddress,
		},
		"invalid recipient": {
			sender:  addrs[0],
			classID: classID,
			err:     sdkerrors.ErrInvalidAddress,
		},
		"invalid class id": {
			sender:    addrs[0],
			recipient: addrs[1],
			err:       internft.ErrInvalidClassID,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			msg := internft.MsgSend{
				Sender:    tc.sender.String(),
				Recipient: tc.recipient.String(),
				Nft: internft.NFT{
					ClassId: tc.classID,
					Id:      math.OneUint(),
				},
			}

			err := msg.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
			if tc.err != nil {
				return
			}

			require.Equal(t, []sdk.AccAddress{tc.sender}, msg.GetSigners())
		})
	}
}

func TestMsgNewClass(t *testing.T) {
	addr := createAddresses(1, "addr")[0]
	const traitID = "uri"

	testCases := map[string]struct {
		owner   sdk.AccAddress
		traitID string
		err     error
	}{
		"valid msg": {
			owner:   addr,
			traitID: traitID,
		},
		"invalid owner": {
			traitID: traitID,
			err:     sdkerrors.ErrInvalidAddress,
		},
		"invalid trait id": {
			owner: addr,
			err:   internft.ErrInvalidTraitID,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			msg := internft.MsgNewClass{
				Owner: tc.owner.String(),
				Traits: []internft.Trait{
					{
						Id: tc.traitID,
					},
				},
			}

			err := msg.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
			if tc.err != nil {
				return
			}

			require.Equal(t, []sdk.AccAddress{tc.owner}, msg.GetSigners())
		})
	}
}

func TestMsgUpdateClass(t *testing.T) {
	classID := createClassIDs(1, "class")[0]

	testCases := map[string]struct {
		classID string
		err     error
	}{
		"valid msg": {
			classID: classID,
		},
		"invalid class id": {
			err: internft.ErrInvalidClassID,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			msg := internft.MsgUpdateClass{
				ClassId: tc.classID,
			}

			err := msg.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
			if tc.err != nil {
				return
			}

			owner := internft.ClassOwner(tc.classID)
			require.Equal(t, []sdk.AccAddress{owner}, msg.GetSigners())
		})
	}
}

func TestMsgMintNFT(t *testing.T) {
	classID := createClassIDs(1, "class")[0]
	traitID := "uri"
	addr := createAddresses(1, "addr")[0]

	testCases := map[string]struct {
		classID   string
		traitID   string
		recipient sdk.AccAddress
		err       error
	}{
		"valid msg": {
			classID:   classID,
			traitID:   traitID,
			recipient: addr,
		},
		"invalid class id": {
			recipient: addr,
			traitID:   traitID,
			err:       internft.ErrInvalidClassID,
		},
		"invalid trait id": {
			classID: classID,
			err:     internft.ErrInvalidTraitID,
		},
		"invalid recipient": {
			classID: classID,
			traitID: traitID,
			err:     sdkerrors.ErrInvalidAddress,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			msg := internft.MsgMintNFT{
				ClassId: tc.classID,
				Properties: []internft.Property{
					{
						Id: tc.traitID,
					},
				},
				Recipient: tc.recipient.String(),
			}

			err := msg.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
			if tc.err != nil {
				return
			}

			owner := internft.ClassOwner(tc.classID)
			require.Equal(t, []sdk.AccAddress{owner}, msg.GetSigners())
		})
	}
}

func TestMsgBurnNFT(t *testing.T) {
	addr := createAddresses(1, "addr")[0]
	classID := createClassIDs(1, "class")[0]

	testCases := map[string]struct {
		owner   sdk.AccAddress
		classID string
		err     error
	}{
		"valid msg": {
			owner:   addr,
			classID: classID,
		},
		"invalid owner": {
			classID: classID,
			err:     sdkerrors.ErrInvalidAddress,
		},
		"invalid class id": {
			owner: addr,
			err:   internft.ErrInvalidClassID,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			msg := internft.MsgBurnNFT{
				Owner: tc.owner.String(),
				Nft: internft.NFT{
					ClassId: tc.classID,
					Id:      math.OneUint(),
				},
			}

			err := msg.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
			if tc.err != nil {
				return
			}

			require.Equal(t, []sdk.AccAddress{tc.owner}, msg.GetSigners())
		})
	}
}

func TestMsgUpdateNFT(t *testing.T) {
	classID := createClassIDs(1, "class")[0]
	traitID := "uri"

	testCases := map[string]struct {
		classID  string
		traitIDs []string
		err      error
	}{
		"valid msg": {
			classID: classID,
			traitIDs: []string{
				traitID,
			},
		},
		"invalid class id": {
			traitIDs: []string{
				traitID,
			},
			err: internft.ErrInvalidClassID,
		},
		"empty properties": {
			classID: classID,
			err:     sdkerrors.ErrInvalidRequest,
		},
		"invalid trait id": {
			classID: classID,
			traitIDs: []string{
				"",
			},
			err: internft.ErrInvalidTraitID,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			properties := make([]internft.Property, len(tc.traitIDs))
			for i, id := range tc.traitIDs {
				properties[i] = internft.Property{
					Id: id,
				}
			}

			msg := internft.MsgUpdateNFT{
				Nft: internft.NFT{
					ClassId: tc.classID,
					Id:      math.OneUint(),
				},
				Properties: properties,
			}

			err := msg.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
			if tc.err != nil {
				return
			}

			owner := internft.ClassOwner(tc.classID)
			require.Equal(t, []sdk.AccAddress{owner}, msg.GetSigners())
		})
	}
}

func TestLegacyMsg(t *testing.T) {
	addrs := createAddresses(2, "addr")
	classIDs := createClassIDs(2, "class")
	id := math.OneUint()
	uri := "https://ipfs.io/ipfs/tIBeTianfOX"

	testCase := []struct {
		msg legacytx.LegacyMsg
		out string
	}{
		{
			&internft.MsgSend{
				Sender:    addrs[0].String(),
				Recipient: addrs[1].String(),
				Nft: internft.NFT{
					ClassId: classIDs[0],
					Id:      id,
				},
			},
			`{"nft":{"class_id":"cosmos1vdkxzumnxq7xavm0","id":"1"},"recipient":"cosmos1v9jxgu33kfsgr5","sender":"cosmos1v9jxgu3stlya7x"}`,
		},
		{
			&internft.MsgNewClass{
				Owner: addrs[0].String(),
				Traits: []internft.Trait{
					{
						Id:      "uri",
						Mutable: true,
					},
				},
			},
			`{"owner":"cosmos1v9jxgu3stlya7x","traits":[{"id":"uri","mutable":true}]}`,
		},
		{
			&internft.MsgUpdateClass{
				ClassId: classIDs[0],
			},
			`{"class_id":"cosmos1vdkxzumnxq7xavm0"}`,
		},
		{
			&internft.MsgMintNFT{
				ClassId: classIDs[0],
				Properties: []internft.Property{
					{
						Id:   "uri",
						Fact: uri,
					},
				},
				Recipient: addrs[0].String(),
			},
			`{"class_id":"cosmos1vdkxzumnxq7xavm0","properties":[{"fact":"https://ipfs.io/ipfs/tIBeTianfOX","id":"uri"}],"recipient":"cosmos1v9jxgu3stlya7x"}`,
		},
		{
			&internft.MsgBurnNFT{
				Owner: addrs[0].String(),
				Nft: internft.NFT{
					ClassId: classIDs[0],
					Id:      id,
				},
			},
			`{"nft":{"class_id":"cosmos1vdkxzumnxq7xavm0","id":"1"},"owner":"cosmos1v9jxgu3stlya7x"}`,
		},
		{
			&internft.MsgUpdateNFT{
				Nft: internft.NFT{
					ClassId: classIDs[0],
					Id:      id,
				},
				Properties: []internft.Property{
					{
						Id:   "uri",
						Fact: uri,
					},
				},
			},
			`{"nft":{"class_id":"cosmos1vdkxzumnxq7xavm0","id":"1"},"properties":[{"fact":"https://ipfs.io/ipfs/tIBeTianfOX","id":"uri"}]}`,
		},
	}

	for _, tc := range testCase {
		name := sdk.MsgTypeURL(tc.msg)[1:]
		t.Run(name, func(t *testing.T) {
			out := tc.msg.GetSignBytes()
			require.Equal(t, tc.out, string(out))
		})
	}
}
