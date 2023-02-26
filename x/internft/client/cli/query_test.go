package cli_test

import (
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"

	"cosmossdk.io/math"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/internft"
	"github.com/cosmos/cosmos-sdk/x/internft/client/cli"
)

func (s *CLITestSuite) TestNewQueryCmdParams() {
	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid request": {
			args: []string{},
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewQueryCmdParams()

			out, err := clitestutil.ExecTestCLICmd(s.clientCtx, cmd, append(tc.args, s.commonQueryFlags...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res internft.QueryParamsResponse
			err = s.clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
		})
	}
}

func (s *CLITestSuite) TestNewQueryCmdClass() {
	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid request": {
			args: []string{
				internft.ClassIDFromOwner(s.vendor),
			},
		},
		"invalid class id": {
			args: []string{
				"",
			},
			err: internft.ErrInvalidClassID,
		},
		// "class not found": {
		// 	args: []string{
		// 		internft.ClassIDFromOwner(s.customer),
		// 	},
		// 	err: status.Error(
		// 		codes.NotFound,
		// 		sdkerrors.ErrKeyNotFound.Wrap(
		// 			status.Error(
		// 				codes.NotFound,
		// 				internft.ErrClassNotFound.Wrap(
		// 					internft.ClassIDFromOwner(s.customer),
		// 				).Error(),
		// 			).Error(),
		// 		).Error(),
		// 	),
		// },
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewQueryCmdClass()

			out, err := clitestutil.ExecTestCLICmd(s.clientCtx, cmd, append(tc.args, s.commonQueryFlags...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res internft.QueryClassResponse
			err = s.clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
		})
	}
}

func (s *CLITestSuite) TestNewQueryCmdClasses() {
	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid request": {
			args: []string{},
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewQueryCmdClasses()

			out, err := clitestutil.ExecTestCLICmd(s.clientCtx, cmd, append(tc.args, s.commonQueryFlags...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res internft.QueryClassesResponse
			err = s.clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
		})
	}
}

func (s *CLITestSuite) TestNewQueryCmdNFT() {
	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid request": {
			args: []string{
				internft.NFT{
					ClassId: internft.ClassIDFromOwner(s.vendor),
					Id:      math.OneUint(),
				}.String(),
			},
		},
		"invalid id": {
			args: []string{
				internft.NFT{
					Id: math.OneUint(),
				}.String(),
			},
			err: internft.ErrInvalidClassID,
		},
		// "nft not found": {
		// 	args: []string{
		// 		internft.NFT{
		// 			ClassId: internft.ClassIDFromOwner(s.customer),
		// 			Id:      math.OneUint(),
		// 		}.String(),
		// 	},
		// 	err: status.Error(
		// 		codes.NotFound,
		// 		sdkerrors.ErrKeyNotFound.Wrap(
		// 			status.Error(
		// 				codes.NotFound,
		// 				internft.ErrNFTNotFound.Wrap(
		// 					(&internft.NFT{
		// 						ClassId: internft.ClassIDFromOwner(s.customer),
		// 						Id:      math.OneUint(),
		// 					}).String(),
		// 				).Error(),
		// 			).Error(),
		// 		).Error(),
		// 	),
		// },
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewQueryCmdNFT()

			out, err := clitestutil.ExecTestCLICmd(s.clientCtx, cmd, append(tc.args, s.commonQueryFlags...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res internft.QueryNFTResponse
			err = s.clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
		})
	}
}

func (s *CLITestSuite) TestNewQueryCmdNFTs() {
	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid request": {
			args: []string{
				internft.ClassIDFromOwner(s.vendor),
			},
		},
		"invalid class id": {
			args: []string{
				"",
			},
			err: internft.ErrInvalidClassID,
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewQueryCmdNFTs()

			out, err := clitestutil.ExecTestCLICmd(s.clientCtx, cmd, append(tc.args, s.commonQueryFlags...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res internft.QueryNFTsResponse
			err = s.clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
		})
	}
}

func (s *CLITestSuite) TestNewQueryCmdOwner() {
	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid request": {
			args: []string{
				internft.NFT{
					ClassId: internft.ClassIDFromOwner(s.vendor),
					Id:      math.OneUint(),
				}.String(),
			},
		},
		"invalid id": {
			args: []string{
				internft.NFT{
					Id: math.OneUint(),
				}.String(),
			},
			err: internft.ErrInvalidClassID,
		},
		// "nft not found": {
		// 	args: []string{
		// 		internft.NFT{
		// 			ClassId: internft.ClassIDFromOwner(s.customer),
		// 			Id:      math.OneUint(),
		// 		}.String(),
		// 	},
		// 	err: status.Error(
		// 		codes.NotFound,
		// 		sdkerrors.ErrKeyNotFound.Wrap(
		// 			status.Error(
		// 				codes.NotFound,
		// 				internft.ErrNFTNotFound.Wrap(
		// 					(&internft.NFT{
		// 						ClassId: internft.ClassIDFromOwner(s.customer),
		// 						Id:      math.OneUint(),
		// 					}).String(),
		// 				).Error(),
		// 			).Error(),
		// 		).Error(),
		// 	),
		// },
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewQueryCmdOwner()

			out, err := clitestutil.ExecTestCLICmd(s.clientCtx, cmd, append(tc.args, s.commonQueryFlags...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res internft.QueryOwnerResponse
			err = s.clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
		})
	}
}
