package keeper

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/internft"
)

func (k Keeper) Send(ctx sdk.Context, sender, recipient sdk.AccAddress, nft internft.NFT) error {
	if err := k.validateOwner(ctx, nft, sender); err != nil {
		return err
	}
	k.setOwner(ctx, nft, recipient)

	return nil
}

func (k Keeper) validateOwner(ctx sdk.Context, nft internft.NFT, owner sdk.AccAddress) error {
	if real, err := k.getOwner(ctx, nft); err != nil || !owner.Equals(real) {
		return errorsmod.Wrap(internft.ErrInsufficientNFT.Wrap("not owns root nft"), nft.String())
	}

	return nil
}

func (k Keeper) GetOwner(ctx sdk.Context, nft internft.NFT) (*sdk.AccAddress, error) {
	if err := k.hasNFT(ctx, nft); err != nil {
		return nil, err
	}

	owner, err := k.getOwner(ctx, nft)
	if err != nil {
		panic(err)
	}

	return owner, nil
}

// func (k Keeper) hasOwner(ctx sdk.Context, nft internft.NFT) error {
// 	_, err := k.getOwnerBytes(ctx, nft)
// 	return err
// }

func (k Keeper) getOwner(ctx sdk.Context, nft internft.NFT) (*sdk.AccAddress, error) {
	bz, err := k.getOwnerBytes(ctx, nft)
	if err != nil {
		return nil, err
	}

	var owner sdk.AccAddress
	if err := owner.Unmarshal(bz); err != nil {
		panic(err)
	}

	return &owner, nil
}

func (k Keeper) getOwnerBytes(ctx sdk.Context, nft internft.NFT) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	key := ownerKey(nft.ClassId, nft.Id)

	bz := store.Get(key)
	if bz == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrNotFound.Wrap("owner"), nft.String())
	}

	return bz, nil
}

func (k Keeper) setOwner(ctx sdk.Context, nft internft.NFT, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	key := ownerKey(nft.ClassId, nft.Id)

	bz, err := owner.Marshal()
	if err != nil {
		panic(err)
	}

	store.Set(key, bz)
}

func (k Keeper) deleteOwner(ctx sdk.Context, nft internft.NFT) {
	store := ctx.KVStore(k.storeKey)
	key := ownerKey(nft.ClassId, nft.Id)

	store.Delete(key)
}
