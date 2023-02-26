package cli

import (
	"encoding/json"

	// "github.com/spf13/cobra"

	errorsmod "cosmossdk.io/errors"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/internft"
)

// func validateGenerateOnly(cmd *cobra.Command) error {
// 	generateOnly, err := cmd.Flags().GetBool(flags.FlagGenerateOnly)
// 	if err != nil {
// 		return err
// 	}

// 	if !generateOnly {
// 		return sdkerrors.ErrNotSupported.Wrapf("must use it with the flag --%s", flags.FlagGenerateOnly)
// 	}

// 	return nil
// }

func ParseParams(codec codec.Codec, paramsJSON string) (*internft.Params, error) {
	var params internft.Params
	if err := codec.UnmarshalJSON([]byte(paramsJSON), &params); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidType.Wrap("params"), err.Error())
	}

	return &params, nil
}

func ParseTraits(codec codec.Codec, traitsJSON string) ([]internft.Trait, error) {
	var traitJSONs []json.RawMessage
	if err := json.Unmarshal([]byte(traitsJSON), &traitJSONs); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidType.Wrap("traits"), err.Error())
	}

	if len(traitJSONs) == 0 {
		return nil, nil
	}

	traits := make([]internft.Trait, len(traitJSONs))
	for i, traitJSON := range traitJSONs {
		if err := codec.UnmarshalJSON(traitJSON, &traits[i]); err != nil {
			return nil, errorsmod.Wrap(sdkerrors.ErrInvalidType.Wrapf("trait %d", i), err.Error())
		}
	}

	return traits, nil
}

func ParseProperties(codec codec.Codec, propertiesJSON string) ([]internft.Property, error) {
	var propertyJSONs []json.RawMessage
	if err := json.Unmarshal([]byte(propertiesJSON), &propertyJSONs); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidType.Wrap("properties"), err.Error())
	}

	if len(propertyJSONs) == 0 {
		return nil, nil
	}

	properties := make([]internft.Property, len(propertyJSONs))
	for i, propertyJSON := range propertyJSONs {
		if err := codec.UnmarshalJSON(propertyJSON, &properties[i]); err != nil {
			return nil, errorsmod.Wrap(sdkerrors.ErrInvalidType.Wrapf("property %d", i), err.Error())
		}
	}

	return properties, nil
}
