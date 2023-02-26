package keeper_test

import (
	"cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/x/internft"
)

func (s *KeeperTestSuite) TestNewClass() {
	testCases := map[string]struct {
		classID string
		err     error
	}{
		"valid request": {
			classID: internft.ClassIDFromOwner(s.customer),
		},
		"class already exists": {
			classID: internft.ClassIDFromOwner(s.vendor),
			err:     internft.ErrClassAlreadyExists,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			class := internft.Class{
				Id: tc.classID,
			}
			err := class.ValidateBasic()
			s.Assert().NoError(err)

			traits := []internft.Trait{
				{
					Id: "uri",
				},
			}

			err = s.keeper.NewClass(ctx, class, traits)
			s.Require().ErrorIs(err, tc.err)
			if err != nil {
				return
			}

			got, err := s.keeper.GetClass(ctx, tc.classID)
			s.Require().NoError(err)
			s.Require().NotNil(got)
			s.Require().Equal(class, *got)
		})
	}
}

func (s *KeeperTestSuite) TestUpdateClass() {
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

			class := internft.Class{
				Id: tc.classID,
			}
			err := class.ValidateBasic()
			s.Assert().NoError(err)

			err = s.keeper.UpdateClass(ctx, class)
			s.Require().ErrorIs(err, tc.err)
			if err != nil {
				return
			}

			got, err := s.keeper.GetClass(ctx, tc.classID)
			s.Require().NoError(err)
			s.Require().NotNil(got)
			s.Require().Equal(class, *got)
		})
	}
}

func (s *KeeperTestSuite) TestMintNFT() {
	testCases := map[string]struct {
		classID    string
		propertyID string
		err        error
	}{
		"valid request": {
			classID:    internft.ClassIDFromOwner(s.vendor),
			propertyID: s.immutableTraitID,
		},
		"class not found": {
			classID:    internft.ClassIDFromOwner(s.customer),
			propertyID: s.immutableTraitID,
			err:        internft.ErrClassNotFound,
		},
		"trait not found": {
			classID:    internft.ClassIDFromOwner(s.vendor),
			propertyID: "no-such-a-trait",
			err:        internft.ErrTraitNotFound,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			err := internft.ValidateClassID(tc.classID)
			s.Assert().NoError(err)

			properties := []internft.Property{
				{
					Id:   tc.propertyID,
					Fact: randomString(32),
				},
			}
			err = internft.Properties(properties).ValidateBasic()
			s.Assert().NoError(err)

			id, err := s.keeper.MintNFT(ctx, s.vendor, tc.classID, properties)
			s.Require().ErrorIs(err, tc.err)
			if err != nil {
				s.Require().Nil(id)
				return
			}
			s.Require().NotNil(id)

			nft := internft.NFT{
				ClassId: tc.classID,
				Id:      *id,
			}
			err = nft.ValidateBasic()
			s.Require().NoError(err)

			_, err = s.keeper.GetNFT(ctx, nft)
			s.Require().NoError(err)

			got, err := s.keeper.GetProperty(ctx, nft, tc.propertyID)
			s.Require().NoError(err)
			s.Require().NotNil(got)
			s.Require().Equal(properties[0], *got)
		})
	}
}

func (s *KeeperTestSuite) TestBurnNFT() {
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

			nft := internft.NFT{
				ClassId: internft.ClassIDFromOwner(s.vendor),
				Id:      tc.id,
			}
			err := nft.ValidateBasic()
			s.Assert().NoError(err)

			err = s.keeper.BurnNFT(ctx, s.vendor, nft)
			s.Require().ErrorIs(err, tc.err)
			if err != nil {
				return
			}

			got, err := s.keeper.GetNFT(ctx, nft)
			s.Require().Error(err)
			s.Require().Nil(got)
		})
	}
}

func (s *KeeperTestSuite) TestUpdateNFT() {
	testCases := map[string]struct {
		id         math.Uint
		propertyID string
		err        error
	}{
		"valid request": {
			id:         math.OneUint(),
			propertyID: s.mutableTraitID,
		},
		"nft not found": {
			id:         math.NewUint(s.numNFTs*2 + 1),
			propertyID: s.mutableTraitID,
			err:        internft.ErrNFTNotFound,
		},
		"trait not found": {
			id:         math.OneUint(),
			propertyID: "no-such-a-trait",
			err:        internft.ErrTraitNotFound,
		},
		"trait immutable": {
			id:         math.OneUint(),
			propertyID: s.immutableTraitID,
			err:        internft.ErrTraitImmutable,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			nft := internft.NFT{
				ClassId: internft.ClassIDFromOwner(s.vendor),
				Id:      tc.id,
			}
			err := nft.ValidateBasic()
			s.Assert().NoError(err)

			property := internft.Property{
				Id:   tc.propertyID,
				Fact: randomString(32),
			}
			err = property.ValidateBasic()
			s.Assert().NoError(err)

			err = s.keeper.UpdateNFT(ctx, nft, []internft.Property{property})
			s.Require().ErrorIs(err, tc.err)
			if err != nil {
				return
			}

			got, err := s.keeper.GetProperty(ctx, nft, tc.propertyID)
			s.Require().NoError(err)
			s.Require().NotNil(got)
			s.Require().Equal(property, *got)
		})
	}
}
