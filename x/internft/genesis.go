package internft

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	math "cosmossdk.io/math"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// DefaultGenesisState - Return a default genesis state
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

func DefaultParams() Params {
	return Params{}
}

// ValidateBasic check the given genesis state has no integrity issues
func (s GenesisState) ValidateBasic() error {
	classIDs := map[string]struct{}{}
	for classIndex, genClass := range s.Classes {
		errHint := fmt.Sprintf("classes[%d]", classIndex)

		id := genClass.Id
		if err := ValidateClassID(id); err != nil {
			return errorsmod.Wrap(err, errHint)
		}

		if _, seen := classIDs[id]; seen {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest.Wrapf("duplicate class id %s", genClass.Id), errHint)
		}
		classIDs[id] = struct{}{}

		if err := Traits(genClass.Traits).ValidateBasic(); err != nil {
			return errorsmod.Wrap(err, errHint)
		}

		traits := map[string]struct{}{}
		for _, trait := range genClass.Traits {
			traits[trait.Id] = struct{}{}
		}

		seenID := math.ZeroUint()
		for nftIndex, genNFT := range genClass.Nfts {
			errHint := fmt.Sprintf("%s.nfts[%d]", errHint, nftIndex)

			id := genNFT.Id
			if err := ValidateNFTID(id); err != nil {
				return errorsmod.Wrap(err, errHint)
			}

			if id.LTE(seenID) {
				return errorsmod.Wrap(sdkerrors.ErrInvalidRequest.Wrap("unsorted nfts"), errHint)
			}
			if id.GT(genClass.LastMintedNftId) {
				return errorsmod.Wrap(sdkerrors.ErrInvalidRequest.Wrapf("id %s > last minted id %s", id, genClass.LastMintedNftId), errHint)
			}
			seenID = id

			if err := Properties(genNFT.Properties).ValidateBasic(); err != nil {
				return errorsmod.Wrap(err, errHint)
			}

			for _, property := range genNFT.Properties {
				if _, hasTrait := traits[property.Id]; !hasTrait {
					return errorsmod.Wrap(ErrTraitNotFound.Wrap(property.Id), errHint)
				}
			}

			if err := ValidateAddress(genNFT.Owner); err != nil {
				return errorsmod.Wrap(errorsmod.Wrap(err, "owner"), errHint)
			}
		}
	}

	return nil
}
