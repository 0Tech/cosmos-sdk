package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/internft"
)

func (k Keeper) NewClass(ctx sdk.Context, class internft.Class, traits []internft.Trait) error {
	if err := k.hasClass(ctx, class.Id); err == nil {
		return internft.ErrClassAlreadyExists.Wrap(class.Id)
	}
	k.setClass(ctx, class)

	for _, trait := range traits {
		k.setTrait(ctx, class.Id, trait)
	}

	k.setPreviousID(ctx, class.Id, math.ZeroUint())

	return nil
}

func (k Keeper) UpdateClass(ctx sdk.Context, class internft.Class) error {
	if err := k.hasClass(ctx, class.Id); err != nil {
		return err
	}
	k.setClass(ctx, class)

	return nil
}

func (k Keeper) hasClass(ctx sdk.Context, classID string) error {
	_, err := k.getClassBytes(ctx, classID)
	return err
}

func (k Keeper) GetClass(ctx sdk.Context, classID string) (*internft.Class, error) {
	bz, err := k.getClassBytes(ctx, classID)
	if err != nil {
		return nil, err
	}

	var class internft.Class
	k.cdc.MustUnmarshal(bz, &class)

	return &class, nil
}

func (k Keeper) getClassBytes(ctx sdk.Context, classID string) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	key := classKey(classID)

	bz := store.Get(key)
	if bz == nil {
		return nil, internft.ErrClassNotFound.Wrap(classID)
	}

	return bz, nil
}

func (k Keeper) setClass(ctx sdk.Context, class internft.Class) {
	store := ctx.KVStore(k.storeKey)
	key := classKey(class.Id)

	bz := k.cdc.MustMarshal(&class)

	store.Set(key, bz)
}

func (k Keeper) hasTrait(ctx sdk.Context, classID string, traitID string) error {
	_, err := k.getTraitBytes(ctx, classID, traitID)
	return err
}

func (k Keeper) GetTrait(ctx sdk.Context, classID string, traitID string) (*internft.Trait, error) {
	bz, err := k.getTraitBytes(ctx, classID, traitID)
	if err != nil {
		return nil, err
	}

	var trait internft.Trait
	k.cdc.MustUnmarshal(bz, &trait)

	return &trait, nil
}

func (k Keeper) getTraitBytes(ctx sdk.Context, classID string, traitID string) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	key := traitKey(classID, traitID)

	bz := store.Get(key)
	if bz == nil {
		return nil, internft.ErrTraitNotFound.Wrapf("%s, %s", classID, traitID)
	}

	return bz, nil
}

func (k Keeper) setTrait(ctx sdk.Context, classID string, trait internft.Trait) {
	store := ctx.KVStore(k.storeKey)
	key := traitKey(classID, trait.Id)

	bz := k.cdc.MustMarshal(&trait)

	store.Set(key, bz)
}

func (k Keeper) GetPreviousID(ctx sdk.Context, classID string) math.Uint {
	bz, err := k.getPreviousIDBytes(ctx, classID)
	if err != nil {
		panic(err)
	}

	var id math.Uint
	if err := id.Unmarshal(bz); err != nil {
		panic(err)
	}

	return id
}

func (k Keeper) getPreviousIDBytes(ctx sdk.Context, classID string) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	key := previousIDKey(classID)

	bz := store.Get(key)
	if bz == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrNotFound.Wrap("previous id"), classID)
	}

	return bz, nil
}

func (k Keeper) setPreviousID(ctx sdk.Context, classID string, id math.Uint) {
	store := ctx.KVStore(k.storeKey)
	key := previousIDKey(classID)

	bz, err := id.Marshal()
	if err != nil {
		panic(err)
	}

	store.Set(key, bz)
}

func (k Keeper) iterateClasses(ctx sdk.Context, fn func(class internft.Class)) {
	prefix := classKeyPrefix

	k.iterateImpl(ctx, prefix, func(_, value []byte) {
		var class internft.Class
		k.cdc.MustUnmarshal(value, &class)

		fn(class)
	})
}

func (k Keeper) MintNFT(ctx sdk.Context, owner sdk.AccAddress, classID string, properties []internft.Property) (*math.Uint, error) {
	if err := k.hasClass(ctx, classID); err != nil {
		return nil, err
	}

	id := k.GetPreviousID(ctx, classID).Incr()
	k.setPreviousID(ctx, classID, id)

	nft := internft.NFT{
		ClassId: classID,
		Id:      id,
	}

	if err := k.hasNFT(ctx, nft); err == nil {
		panic(errorsmod.Wrap(sdkerrors.ErrInvalidRequest.Wrap("nft already exists"), nft.String()))
	}
	k.setNFT(ctx, nft)

	for _, property := range properties {
		if err := k.hasTrait(ctx, nft.ClassId, property.Id); err != nil {
			return nil, errorsmod.Wrap(err, property.Id)
		}

		k.setProperty(ctx, nft, property)
	}

	k.setOwner(ctx, nft, owner)

	return &nft.Id, nil
}

// func (k Keeper) hasProperty(ctx sdk.Context, nft internft.NFT, propertyID string) error {
// 	_, err := k.getPropertyBytes(ctx, nft, propertyID)
// 	return err
// }

func (k Keeper) GetProperty(ctx sdk.Context, nft internft.NFT, propertyID string) (*internft.Property, error) {
	bz, err := k.getPropertyBytes(ctx, nft, propertyID)
	if err != nil {
		return nil, err
	}

	var property internft.Property
	k.cdc.MustUnmarshal(bz, &property)

	return &property, nil
}

func (k Keeper) getPropertyBytes(ctx sdk.Context, nft internft.NFT, propertyID string) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	key := propertyKey(nft.ClassId, nft.Id, propertyID)

	bz := store.Get(key)
	if bz == nil {
		return nil, internft.ErrTraitNotFound.Wrapf("%s, %s", nft.ClassId, propertyID)
	}

	return bz, nil
}

func (k Keeper) setProperty(ctx sdk.Context, nft internft.NFT, property internft.Property) {
	store := ctx.KVStore(k.storeKey)
	key := propertyKey(nft.ClassId, nft.Id, property.Id)

	bz := k.cdc.MustMarshal(&property)

	store.Set(key, bz)
}

func (k Keeper) BurnNFT(ctx sdk.Context, owner sdk.AccAddress, nft internft.NFT) error {
	if err := k.validateOwner(ctx, nft, owner); err != nil {
		return err
	}
	k.deleteOwner(ctx, nft)

	if err := k.hasNFT(ctx, nft); err != nil {
		panic(err)
	}
	k.deleteNFT(ctx, nft)

	// TODO: prune children

	return nil
}

func (k Keeper) UpdateNFT(ctx sdk.Context, nft internft.NFT, properties []internft.Property) error {
	if err := k.hasNFT(ctx, nft); err != nil {
		return err
	}

	for _, property := range properties {
		trait, err := k.GetTrait(ctx, nft.ClassId, property.Id)
		if err != nil {
			return err
		}

		if !trait.Mutable {
			return internft.ErrTraitImmutable.Wrap(property.Id)
		}

		k.setProperty(ctx, nft, property)
	}

	return nil
}

func (k Keeper) hasNFT(ctx sdk.Context, nft internft.NFT) error {
	_, err := k.getNFTBytes(ctx, nft)
	return err
}

func (k Keeper) GetNFT(ctx sdk.Context, nft internft.NFT) (*internft.NFT, error) {
	if err := k.hasNFT(ctx, nft); err != nil {
		return nil, err
	}

	return &nft, nil
}

func (k Keeper) getNFTBytes(ctx sdk.Context, nft internft.NFT) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	key := nftKey(nft.ClassId, nft.Id)

	bz := store.Get(key)
	if bz == nil {
		return nil, internft.ErrNFTNotFound.Wrap(nft.String())
	}

	return bz, nil
}

func (k Keeper) setNFT(ctx sdk.Context, nft internft.NFT) {
	store := ctx.KVStore(k.storeKey)
	key := nftKey(nft.ClassId, nft.Id)

	bz := k.cdc.MustMarshal(&nft)

	store.Set(key, bz)
}

func (k Keeper) deleteNFT(ctx sdk.Context, nft internft.NFT) {
	store := ctx.KVStore(k.storeKey)
	key := nftKey(nft.ClassId, nft.Id)

	store.Delete(key)
}

func (k Keeper) iterateNFTsOfClass(ctx sdk.Context, classID string, fn func(nft internft.NFT)) {
	prefix := nftKeyPrefixOfClass(classID)

	k.iterateImpl(ctx, prefix, func(_, value []byte) {
		var nft internft.NFT
		k.cdc.MustUnmarshal(value, &nft)

		fn(nft)
	})
}
