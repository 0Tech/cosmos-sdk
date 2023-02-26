package keeper_test

import (
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/internft"
)

func (s *KeeperTestSuite) TestSend() {
	testCases := map[string]struct {
		sender sdk.AccAddress
		id     math.Uint
		err    error
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

			err = s.keeper.Send(ctx, s.vendor, s.customer, nft)
			s.Require().ErrorIs(err, tc.err)
			if err != nil {
				return
			}

			got, err := s.keeper.GetOwner(ctx, nft)
			s.Require().NoError(err)
			s.Require().Equal(s.customer, *got)
		})
	}
}
