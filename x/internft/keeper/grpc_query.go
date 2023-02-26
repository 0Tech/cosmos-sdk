package keeper

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmossdk.io/store/prefix"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/internft"
)

type queryServer struct {
	keeper Keeper
}

var _ internft.QueryServer = (*queryServer)(nil)

// NewQueryServer returns an implementation of the QueryServer interface
// for the provided Keeper.
func NewQueryServer(keeper Keeper) internft.QueryServer {
	return &queryServer{
		keeper: keeper,
	}
}

const didDelimiter = ":"

// Params queries the module params.
func (s queryServer) Params(c context.Context, req *internft.QueryParamsRequest) (*internft.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	params := s.keeper.GetParams(ctx)

	return &internft.QueryParamsResponse{
		Params: params,
	}, nil
}

// Class queries a class.
func (s queryServer) Class(c context.Context, req *internft.QueryClassRequest) (*internft.QueryClassResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	if err := internft.ValidateClassID(req.ClassId); err != nil {
		return nil, err
	}

	class, err := s.keeper.GetClass(ctx, req.ClassId)
	if err != nil {
		return nil, err
	}

	return &internft.QueryClassResponse{
		Class: class,
	}, nil
}

// Classes queries all classes.
func (s queryServer) Classes(c context.Context, req *internft.QueryClassesRequest) (*internft.QueryClassesResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(s.keeper.storeKey)
	classStore := prefix.NewStore(store, classKeyPrefix)

	var classes []internft.Class
	pageRes, err := query.Paginate(classStore, req.Pagination, func(_ []byte, value []byte) error {
		var class internft.Class
		s.keeper.cdc.MustUnmarshal(value, &class)

		classes = append(classes, class)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &internft.QueryClassesResponse{
		Classes:    classes,
		Pagination: pageRes,
	}, nil
}

// Trait queries a trait of a class.
func (s queryServer) Trait(c context.Context, req *internft.QueryTraitRequest) (*internft.QueryTraitResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	if err := internft.ValidateClassID(req.ClassId); err != nil {
		return nil, err
	}

	if err := internft.ValidateTraitID(req.TraitId); err != nil {
		return nil, err
	}

	trait, err := s.keeper.GetTrait(ctx, req.ClassId, req.TraitId)
	if err != nil {
		return nil, err
	}

	return &internft.QueryTraitResponse{
		Trait: trait,
	}, nil
}

// Traits queries all traits of a class.
func (s queryServer) Traits(c context.Context, req *internft.QueryTraitsRequest) (*internft.QueryTraitsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	if err := internft.ValidateClassID(req.ClassId); err != nil {
		return nil, err
	}

	store := ctx.KVStore(s.keeper.storeKey)
	traitStore := prefix.NewStore(store, traitKeyPrefixOfClass(req.ClassId))

	var traits []internft.Trait
	pageRes, err := query.Paginate(traitStore, req.Pagination, func(_ []byte, value []byte) error {
		var trait internft.Trait
		s.keeper.cdc.MustUnmarshal(value, &trait)

		traits = append(traits, trait)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &internft.QueryTraitsResponse{
		Traits:     traits,
		Pagination: pageRes,
	}, nil
}

// NFT queries an nft.
func (s queryServer) NFT(c context.Context, req *internft.QueryNFTRequest) (*internft.QueryNFTResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	did := strings.Join([]string{
		req.ClassId,
		req.Id,
	}, didDelimiter)

	nft, err := internft.NFTFromString(did)
	if err != nil {
		return nil, err
	}

	if err := s.keeper.hasNFT(ctx, *nft); err != nil {
		return nil, err
	}

	return &internft.QueryNFTResponse{
		Nft: nft,
	}, nil
}

// NFTs queries all nfts.
func (s queryServer) NFTs(c context.Context, req *internft.QueryNFTsRequest) (*internft.QueryNFTsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	if err := internft.ValidateClassID(req.ClassId); err != nil {
		return nil, err
	}

	store := ctx.KVStore(s.keeper.storeKey)
	nftStore := prefix.NewStore(store, nftKeyPrefixOfClass(req.ClassId))

	var nfts []internft.NFT
	pageRes, err := query.Paginate(nftStore, req.Pagination, func(_ []byte, value []byte) error {
		var nft internft.NFT
		s.keeper.cdc.MustUnmarshal(value, &nft)

		nfts = append(nfts, nft)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &internft.QueryNFTsResponse{
		Nfts:       nfts,
		Pagination: pageRes,
	}, nil
}

// Property queries a property of a class.
func (s queryServer) Property(c context.Context, req *internft.QueryPropertyRequest) (*internft.QueryPropertyResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	did := strings.Join([]string{
		req.ClassId,
		req.Id,
	}, didDelimiter)

	nft, err := internft.NFTFromString(did)
	if err != nil {
		return nil, err
	}

	if err := internft.ValidateTraitID(req.PropertyId); err != nil {
		return nil, err
	}

	property, err := s.keeper.GetProperty(ctx, *nft, req.PropertyId)
	if err != nil {
		return nil, err
	}

	return &internft.QueryPropertyResponse{
		Property: property,
	}, nil
}

// Properties queries all properties of a class.
func (s queryServer) Properties(c context.Context, req *internft.QueryPropertiesRequest) (*internft.QueryPropertiesResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	did := strings.Join([]string{
		req.ClassId,
		req.Id,
	}, didDelimiter)

	nft, err := internft.NFTFromString(did)
	if err != nil {
		return nil, err
	}

	store := ctx.KVStore(s.keeper.storeKey)
	propertyStore := prefix.NewStore(store, propertyKeyPrefixOfNFT(nft.ClassId, nft.Id))

	var properties []internft.Property
	pageRes, err := query.Paginate(propertyStore, req.Pagination, func(_ []byte, value []byte) error {
		var property internft.Property
		s.keeper.cdc.MustUnmarshal(value, &property)

		properties = append(properties, property)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &internft.QueryPropertiesResponse{
		Properties: properties,
		Pagination: pageRes,
	}, nil
}

// Owner queries the owner of an nft.
func (s queryServer) Owner(c context.Context, req *internft.QueryOwnerRequest) (*internft.QueryOwnerResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	did := strings.Join([]string{
		req.ClassId,
		req.Id,
	}, didDelimiter)

	nft, err := internft.NFTFromString(did)
	if err != nil {
		return nil, err
	}

	owner, err := s.keeper.GetOwner(ctx, *nft)
	if err != nil {
		return nil, err
	}

	return &internft.QueryOwnerResponse{
		Owner: owner.String(),
	}, nil
}
