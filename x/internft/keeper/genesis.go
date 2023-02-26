package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/internft"
)

func (k Keeper) InitGenesis(ctx sdk.Context, gs *internft.GenesisState) error {
	k.SetParams(ctx, gs.Params)

	for _, genClass := range gs.Classes {
		class := internft.Class{
			Id: genClass.Id,
		}
		k.setClass(ctx, class)

		// TODO: trait

		k.setPreviousID(ctx, class.Id, genClass.LastMintedNftId)

		for _, genNFT := range genClass.Nfts {
			nft := internft.NFT{
				ClassId: class.Id,
				Id:      genNFT.Id,
			}
			k.setNFT(ctx, nft)

			// TODO: property

			owner := genNFT.Owner
			k.setOwner(ctx, nft, sdk.MustAccAddressFromBech32(owner))
		}
	}

	return nil
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *internft.GenesisState {
	classes := k.getClasses(ctx)

	var genClasses []internft.GenesisClass
	if len(classes) != 0 {
		genClasses = make([]internft.GenesisClass, len(classes))
	}

	for classIndex, class := range classes {
		genClasses[classIndex].Id = class.Id
		genClasses[classIndex].LastMintedNftId = k.GetPreviousID(ctx, class.Id)

		// TODO: trait

		nfts := k.getNFTsOfClass(ctx, class.Id)

		var genNFTs []internft.GenesisNFT
		if len(nfts) != 0 {
			genNFTs = make([]internft.GenesisNFT, len(nfts))
		}

		for nftIndex, nft := range nfts {
			genNFTs[nftIndex].Id = nft.Id

			// TODO: property

			owner, err := k.getOwner(ctx, nft)
			if err != nil {
				panic(err)
			}
			genNFTs[nftIndex].Owner = owner.String()
		}

		genClasses[classIndex].Nfts = genNFTs
	}

	return &internft.GenesisState{
		Params:  k.GetParams(ctx),
		Classes: genClasses,
	}
}

func (k Keeper) getClasses(ctx sdk.Context) (classes []internft.Class) {
	k.iterateClasses(ctx, func(class internft.Class) {
		classes = append(classes, class)
	})

	return
}

func (k Keeper) getNFTsOfClass(ctx sdk.Context, classID string) (nfts []internft.NFT) {
	k.iterateNFTsOfClass(ctx, classID, func(nft internft.NFT) {
		nfts = append(nfts, nft)
	})

	return
}
