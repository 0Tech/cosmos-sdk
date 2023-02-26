package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/internft"
)

type msgServer struct {
	keeper Keeper
}

var _ internft.MsgServer = (*msgServer)(nil)

// NewMsgServer returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServer(keeper Keeper) internft.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

// Send defines a method to send an nft from one account to another account.
func (s msgServer) Send(c context.Context, req *internft.MsgSend) (*internft.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	sender := sdk.MustAccAddressFromBech32(req.Sender)
	recipient := sdk.MustAccAddressFromBech32(req.Recipient)

	if err := s.keeper.Send(ctx, sender, recipient, req.Nft); err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&internft.EventSend{
		Sender:   req.Sender,
		Receiver: req.Recipient,
		Nft:      req.Nft,
	}); err != nil {
		panic(err)
	}

	return &internft.MsgSendResponse{}, nil
}

// NewClass defines a method to create a class.
func (s msgServer) NewClass(c context.Context, req *internft.MsgNewClass) (*internft.MsgNewClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	owner := sdk.MustAccAddressFromBech32(req.Owner)
	id := internft.ClassIDFromOwner(owner)
	class := internft.Class{
		Id: id,
	}

	if err := s.keeper.NewClass(ctx, class, req.Traits); err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&internft.EventNewClass{
		Class:  class,
		Traits: req.Traits, // TODO: sort
		Data:   req.Data,
	}); err != nil {
		panic(err)
	}

	return &internft.MsgNewClassResponse{}, nil
}

// UpdateClass defines a method to update a class.
func (s msgServer) UpdateClass(c context.Context, req *internft.MsgUpdateClass) (*internft.MsgUpdateClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	class := internft.Class{
		Id: req.ClassId,
	}

	// TODO: data

	if err := s.keeper.UpdateClass(ctx, class); err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&internft.EventUpdateClass{
		Class: class,
		Data:  req.Data,
	}); err != nil {
		panic(err)
	}

	return &internft.MsgUpdateClassResponse{}, nil
}

// MintNFT defines a method to mint an nft.
func (s msgServer) MintNFT(c context.Context, req *internft.MsgMintNFT) (*internft.MsgMintNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	recipient := sdk.MustAccAddressFromBech32(req.Recipient)

	id, err := s.keeper.MintNFT(ctx, recipient, req.ClassId, req.Properties)
	if err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&internft.EventMintNFT{
		Nft: internft.NFT{
			ClassId: req.ClassId,
			Id:      *id,
		},
		Properties: req.Properties, // TODO: sort
		Recipient:  req.Recipient,
	}); err != nil {
		panic(err)
	}

	return &internft.MsgMintNFTResponse{}, nil
}

// BurnNFT defines a method to burn an nft.
func (s msgServer) BurnNFT(c context.Context, req *internft.MsgBurnNFT) (*internft.MsgBurnNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	owner := sdk.MustAccAddressFromBech32(req.Owner)

	if err := s.keeper.BurnNFT(ctx, owner, req.Nft); err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&internft.EventBurnNFT{
		Owner: req.Owner,
		Nft:   req.Nft,
	}); err != nil {
		panic(err)
	}

	return &internft.MsgBurnNFTResponse{}, nil
}

// UpdateNFT defines a method to update an nft.
func (s msgServer) UpdateNFT(c context.Context, req *internft.MsgUpdateNFT) (*internft.MsgUpdateNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if err := s.keeper.UpdateNFT(ctx, req.Nft, req.Properties); err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&internft.EventUpdateNFT{
		Nft:        req.Nft,
		Properties: req.Properties,
	}); err != nil {
		panic(err)
	}

	return &internft.MsgUpdateNFTResponse{}, nil
}
