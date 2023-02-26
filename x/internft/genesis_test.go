package internft_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/internft"
)

func TestGenesisState(t *testing.T) {
	classIDs := createClassIDs(2, "class")
	const traitID = "uri"
	addr := createAddresses(1, "addr")[0]

	testCases := map[string]struct {
		s   internft.GenesisState
		err error
	}{
		"default genesis": {
			s: *internft.DefaultGenesisState(),
		},
		"all features": {
			s: internft.GenesisState{
				Params: internft.DefaultParams(),
				Classes: []internft.GenesisClass{
					{
						Id: classIDs[0],
						Traits: []internft.Trait{
							{
								Id: traitID,
							},
						},
						LastMintedNftId: math.NewUint(2),
						Nfts: []internft.GenesisNFT{
							{
								Id:    math.NewUint(1),
								Owner: addr.String(),
							},
							{
								Id:    math.NewUint(2),
								Owner: addr.String(),
							},
						},
					},
					{
						Id:              classIDs[1],
						LastMintedNftId: math.NewUint(2),
						Nfts: []internft.GenesisNFT{
							{
								Id:    math.NewUint(1),
								Owner: addr.String(),
							},
							{
								Id:    math.NewUint(2),
								Owner: addr.String(),
							},
						},
					},
				},
			},
		},
		"invalid class id": {
			s: internft.GenesisState{
				Params: internft.DefaultParams(),
				Classes: []internft.GenesisClass{
					{
						LastMintedNftId: math.NewUint(2),
					},
				},
			},
			err: internft.ErrInvalidClassID,
		},
		"invalid trait id": {
			s: internft.GenesisState{
				Params: internft.DefaultParams(),
				Classes: []internft.GenesisClass{
					{
						Id: classIDs[0],
						Traits: []internft.Trait{
							{},
						},
						LastMintedNftId: math.NewUint(2),
					},
				},
			},
			err: internft.ErrInvalidTraitID,
		},
		"duplicate class": {
			s: internft.GenesisState{
				Params: internft.DefaultParams(),
				Classes: []internft.GenesisClass{
					{
						Id:              classIDs[0],
						LastMintedNftId: math.NewUint(2),
					},
					{
						Id:              classIDs[0],
						LastMintedNftId: math.NewUint(2),
					},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		"invalid nft id": {
			s: internft.GenesisState{
				Params: internft.DefaultParams(),
				Classes: []internft.GenesisClass{
					{
						Id:              classIDs[0],
						LastMintedNftId: math.NewUint(2),
						Nfts: []internft.GenesisNFT{
							{
								Id:    math.NewUint(0),
								Owner: addr.String(),
							},
						},
					},
				},
			},
			err: internft.ErrInvalidNFTID,
		},
		"unsorted nfts": {
			s: internft.GenesisState{
				Params: internft.DefaultParams(),
				Classes: []internft.GenesisClass{
					{
						Id:              classIDs[0],
						LastMintedNftId: math.NewUint(2),
						Nfts: []internft.GenesisNFT{
							{
								Id:    math.NewUint(2),
								Owner: addr.String(),
							},
							{
								Id:    math.NewUint(1),
								Owner: addr.String(),
							},
						},
					},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		"greater than last minted nft id": {
			s: internft.GenesisState{
				Params: internft.DefaultParams(),
				Classes: []internft.GenesisClass{
					{
						Id:              classIDs[0],
						LastMintedNftId: math.NewUint(0),
						Nfts: []internft.GenesisNFT{
							{
								Id:    math.NewUint(1),
								Owner: addr.String(),
							},
						},
					},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		"invalid property id": {
			s: internft.GenesisState{
				Params: internft.DefaultParams(),
				Classes: []internft.GenesisClass{
					{
						Id: classIDs[0],
						Traits: []internft.Trait{
							{
								Id: traitID,
							},
						},
						LastMintedNftId: math.NewUint(2),
						Nfts: []internft.GenesisNFT{
							{
								Id: math.NewUint(1),
								Properties: []internft.Property{
									{},
								},
								Owner: addr.String(),
							},
						},
					},
				},
			},
			err: internft.ErrInvalidTraitID,
		},
		"no corresponding trait": {
			s: internft.GenesisState{
				Params: internft.DefaultParams(),
				Classes: []internft.GenesisClass{
					{
						Id: classIDs[0],
						Traits: []internft.Trait{
							{
								Id: traitID,
							},
						},
						LastMintedNftId: math.NewUint(2),
						Nfts: []internft.GenesisNFT{
							{
								Id: math.NewUint(1),
								Properties: []internft.Property{
									{
										Id: "nosuchid",
									},
								},
								Owner: addr.String(),
							},
						},
					},
				},
			},
			err: internft.ErrTraitNotFound,
		},
		"invalid owner": {
			s: internft.GenesisState{
				Params: internft.DefaultParams(),
				Classes: []internft.GenesisClass{
					{
						Id:              classIDs[0],
						LastMintedNftId: math.NewUint(2),
						Nfts: []internft.GenesisNFT{
							{
								Id:    math.NewUint(1),
								Owner: "invalid",
							},
						},
					},
				},
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.s.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
		})
	}
}
