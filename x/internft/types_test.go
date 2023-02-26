package internft_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/internft"
)

func createAddresses(size int, prefix string) []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, size)
	for i := range addrs {
		addrs[i] = sdk.AccAddress(fmt.Sprintf("%s%d", prefix, i))
	}

	return addrs
}

func createClassIDs(size int, prefix string) []string {
	owners := createAddresses(size, prefix)
	ids := make([]string, len(owners))
	for i, owner := range owners {
		ids[i] = internft.ClassIDFromOwner(owner)
	}

	return ids
}

func TestClass(t *testing.T) {
	id := createClassIDs(1, "class")[0]

	testCases := map[string]struct {
		id  string
		err error
	}{
		"valid class": {
			id: id,
		},
		"invalid id": {
			err: internft.ErrInvalidClassID,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			class := internft.Class{
				Id: tc.id,
			}

			err := class.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestTraits(t *testing.T) {
	testCases := map[string]struct {
		ids []string
		err error
	}{
		"valid traits": {},
		"invalid id": {
			ids: []string{
				"",
			},
			err: internft.ErrInvalidTraitID,
		},
		"duplicate id": {
			ids: []string{
				"uri",
				"uri",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			traits := make([]internft.Trait, len(tc.ids))
			for i, id := range tc.ids {
				traits[i] = internft.Trait{
					Id: id,
				}
			}

			err := internft.Traits(traits).ValidateBasic()
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestNFT(t *testing.T) {
	classIDs := createClassIDs(2, "class")

	// ValidateBasic()
	testCases := map[string]struct {
		classID string
		id      math.Uint
		err     error
	}{
		"valid nft": {
			classID: classIDs[0],
			id:      math.OneUint(),
		},
		"invalid class id": {
			id:  math.OneUint(),
			err: internft.ErrInvalidClassID,
		},
		"invalid id": {
			classID: classIDs[0],
			id:      math.ZeroUint(),
			err:     internft.ErrInvalidNFTID,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			nft := internft.NFT{
				ClassId: tc.classID,
				Id:      tc.id,
			}

			err := nft.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
		})
	}

	l := internft.NFT{
		ClassId: classIDs[0],
		Id:      math.OneUint(),
	}

	// Equal()
	testCases2 := map[string]struct {
		classID string
		id      math.Uint
		equals  bool
	}{
		"equals": {
			classID: l.ClassId,
			id:      l.Id,
			equals:  true,
		},
		"different class id": {
			classID: classIDs[1],
			id:      l.Id,
		},
		"different id": {
			classID: l.ClassId,
			id:      l.Id.Incr(),
		},
	}

	for name, tc := range testCases2 {
		t.Run(name, func(t *testing.T) {
			r := internft.NFT{
				ClassId: tc.classID,
				Id:      tc.id,
			}

			require.NoError(t, l.ValidateBasic())
			require.NoError(t, r.ValidateBasic())
			require.Equal(t, tc.equals, l.Equal(r))
		})
	}

	bigID := make([]rune, 78)
	for i := range bigID {
		bigID[i] = '0'
	}
	bigID[0] = '1'

	// NFTFromDID
	testCases3 := map[string]struct {
		classID   string
		delimiter string
		id        string
		err       error
	}{
		"valid did": {
			classID:   classIDs[0],
			delimiter: ":",
			id:        string(bigID),
		},
		"invalid format": {
			classID: classIDs[0],
			id:      string(bigID),
			err:     sdkerrors.ErrInvalidType,
		},
		"invalid uint": {
			classID:   classIDs[0],
			delimiter: ":",
			id:        string(bigID) + "0",
			err:       internft.ErrInvalidNFTID,
		},
		"invalid class id": {
			delimiter: ":",
			id:        string(bigID),
			err:       internft.ErrInvalidClassID,
		},
	}

	for name, tc := range testCases3 {
		t.Run(name, func(t *testing.T) {
			did := fmt.Sprintf("%s%s%s", tc.classID, tc.delimiter, tc.id)

			nft, err := internft.NFTFromString(did)
			require.ErrorIs(t, err, tc.err)
			if tc.err != nil {
				return
			}

			require.Equal(t, tc.classID, nft.ClassId)
			require.Equal(t, math.NewUintFromString(tc.id), nft.Id)
		})
	}
}

func TestProperties(t *testing.T) {
	testCases := map[string]struct {
		ids []string
		err error
	}{
		"valid properties": {},
		"invalid id": {
			ids: []string{
				"",
			},
			err: internft.ErrInvalidTraitID,
		},
		"duplicate id": {
			ids: []string{
				"uri",
				"uri",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			properties := make([]internft.Property, len(tc.ids))
			for i, id := range tc.ids {
				properties[i] = internft.Property{
					Id: id,
				}
			}

			err := internft.Properties(properties).ValidateBasic()
			require.ErrorIs(t, err, tc.err)
		})
	}
}
