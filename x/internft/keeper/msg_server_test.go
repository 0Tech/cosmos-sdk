package keeper_test

import (
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/internft"
)

func (s *KeeperTestSuite) TestMsgSend() {
	testCases := map[string]struct {
		id  math.Uint
		err error
	}{
		"valid request": {
			id: math.OneUint(),
		},
		"insufficient funds": {
			id:  math.NewUint(s.numNFTs + 1),
			err: internft.ErrInsufficientNFT,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &internft.MsgSend{
				Sender:    s.vendor.String(),
				Recipient: s.customer.String(),
				Nft: internft.NFT{
					ClassId: internft.ClassIDFromOwner(s.vendor),
					Id:      tc.id,
				},
			}
			err := req.ValidateBasic()
			s.Assert().NoError(err)

			res, err := s.msgServer.Send(ctx, req)
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}
			s.Require().NotNil(res)
		})
	}
}

func (s *KeeperTestSuite) TestMsgNewClass() {
	testCases := map[string]struct {
		owner sdk.AccAddress
		err   error
	}{
		"valid request": {
			owner: s.customer,
		},
		"class already exists": {
			owner: s.vendor,
			err:   internft.ErrClassAlreadyExists,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &internft.MsgNewClass{
				Owner: tc.owner.String(),
			}
			err := req.ValidateBasic()
			s.Assert().NoError(err)

			res, err := s.msgServer.NewClass(ctx, req)
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}
			s.Require().NotNil(res)
		})
	}
}

func (s *KeeperTestSuite) TestMsgUpdateClass() {
	testCases := map[string]struct {
		classID string
		err     error
	}{
		"valid request": {
			classID: internft.ClassIDFromOwner(s.vendor),
		},
		"class not found": {
			classID: internft.ClassIDFromOwner(s.customer),
			err:     internft.ErrClassNotFound,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &internft.MsgUpdateClass{
				ClassId: tc.classID,
			}
			err := req.ValidateBasic()
			s.Assert().NoError(err)

			res, err := s.msgServer.UpdateClass(ctx, req)
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}
			s.Require().NotNil(res)
		})
	}
}

func (s *KeeperTestSuite) TestMsgMintNFT() {
	testCases := map[string]struct {
		classID string
		err     error
	}{
		"valid request": {
			classID: internft.ClassIDFromOwner(s.vendor),
		},
		"class not found": {
			classID: internft.ClassIDFromOwner(s.customer),
			err:     internft.ErrClassNotFound,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &internft.MsgMintNFT{
				ClassId: tc.classID,
				Properties: []internft.Property{
					{
						Id: s.mutableTraitID,
					},
				},
				Recipient: s.customer.String(),
			}
			err := req.ValidateBasic()
			s.Assert().NoError(err)

			res, err := s.msgServer.MintNFT(ctx, req)
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}
			s.Require().NotNil(res)
		})
	}
}

func (s *KeeperTestSuite) TestMsgBurnNFT() {
	testCases := map[string]struct {
		id  math.Uint
		err error
	}{
		"valid request": {
			id: math.OneUint(),
		},
		"insufficient nft": {
			id:  math.NewUint(s.numNFTs + 1),
			err: internft.ErrInsufficientNFT,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &internft.MsgBurnNFT{
				Owner: s.vendor.String(),
				Nft: internft.NFT{
					ClassId: internft.ClassIDFromOwner(s.vendor),
					Id:      tc.id,
				},
			}
			err := req.ValidateBasic()
			s.Assert().NoError(err)

			res, err := s.msgServer.BurnNFT(ctx, req)
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}
			s.Require().NotNil(res)
		})
	}
}

func (s *KeeperTestSuite) TestMsgUpdateNFT() {
	testCases := map[string]struct {
		id  math.Uint
		err error
	}{
		"valid request": {
			id: math.OneUint(),
		},
		"nft not found": {
			id:  math.NewUint(s.numNFTs*2 + 1),
			err: internft.ErrNFTNotFound,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &internft.MsgUpdateNFT{
				Nft: internft.NFT{
					ClassId: internft.ClassIDFromOwner(s.vendor),
					Id:      tc.id,
				},
				Properties: []internft.Property{
					{
						Id: s.mutableTraitID,
					},
				},
			}
			err := req.ValidateBasic()
			s.Assert().NoError(err)

			res, err := s.msgServer.UpdateNFT(ctx, req)
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}
			s.Require().NotNil(res)
		})
	}
}
