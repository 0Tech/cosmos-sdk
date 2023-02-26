package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/internft"
)

func (k Keeper) GetParams(ctx sdk.Context) internft.Params {
	bz, err := k.getParamsBytes(ctx)
	if err != nil {
		panic(err)
	}

	var params internft.Params
	k.cdc.MustUnmarshal(bz, &params)

	return params
}

func (k Keeper) getParamsBytes(ctx sdk.Context) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	key := paramsKey

	bz := store.Get(key)
	if bz == nil {
		return nil, sdkerrors.ErrNotFound.Wrap("params")
	}

	return bz, nil
}

func (k Keeper) SetParams(ctx sdk.Context, params internft.Params) {
	store := ctx.KVStore(k.storeKey)
	key := paramsKey

	bz := k.cdc.MustMarshal(&params)

	store.Set(key, bz)
}
