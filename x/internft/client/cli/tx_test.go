package cli_test

import (
	// "fmt"

	"cosmossdk.io/math"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	// txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/x/internft"
	"github.com/cosmos/cosmos-sdk/x/internft/client/cli"
)

// func (s *CLITestSuite) TestNewTxCmdUpdateParams() {
// 	commonFlags := []string{
// 		fmt.Sprintf("--%s", flags.FlagGenerateOnly),
// 	}

// 	testCases := map[string]struct {
// 		args  []string
// 		err error
// 	}{
// 		"valid request": {
// 			args: []string{
// 				s.authority.String(),
// 				fmt.Sprint(`{}`),
// 			},
// 		},
// 	}

// 	for name, tc := range testCases {
// 		tc := tc

// 		s.Run(name, func() {
// 			cmd := cli.NewTxCmdUpdateParams()

// 			out, err := clitestutil.ExecTestCLICmd(s.clientCtx, cmd, append(tc.args, commonFlags...))
// 			s.Require().ErrorIs(err, tc.err)
// 			if tc.err != nil  {
// 				return
// 			}

// 			var res sdk.TxResponse
// 			err = s.clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
// 			s.Require().NoError(err, out)
// 			s.Require().NotZero(res.Code, out)
// 		})
// 	}
// }

func (s *CLITestSuite) TestNewTxCmdSend() {
	testCases := map[string]struct {
		args []string
		err  error
	}{
		// "valid failing request": {
		// 	args: []string{
		// 		s.customer.String(),
		// 		s.vendor.String(),
		// 		internft.NFT{
		// 			ClassId: internft.ClassIDFromOwner(s.customer),
		// 			Id:      math.OneUint(),
		// 		}.String(),
		// 	},
		// },
		"invalid id": {
			args: []string{
				s.customer.String(),
				s.vendor.String(),
				"",
			},
			err: sdkerrors.ErrInvalidType,
		},
		"invalid msg": {
			args: []string{
				"",
				s.vendor.String(),
				internft.NFT{
					ClassId: internft.ClassIDFromOwner(s.customer),
					Id:      math.OneUint(),
				}.String(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewTxCmdSend()

			out, err := clitestutil.ExecTestCLICmd(s.clientCtx, cmd, append(tc.args, s.commonTxFlags...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res sdk.TxResponse
			err = s.clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
			s.Require().NotZero(res.Code, out)
		})
	}
}

func (s *CLITestSuite) TestNewTxCmdNewClass() {
	testCases := map[string]struct {
		args []string
		err  error
	}{
		// "valid failing request": {
		// 	args: []string{
		// 		s.vendor.String(),
		// 		"[]",
		// 	},
		// },
		"invalid traits": {
			args: []string{
				s.vendor.String(),
				"",
			},
			err: sdkerrors.ErrInvalidType,
		},
		"invalid msg": {
			args: []string{
				"",
				"[]",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewTxCmdNewClass()

			out, err := clitestutil.ExecTestCLICmd(s.clientCtx, cmd, append(tc.args, s.commonTxFlags...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res sdk.TxResponse
			err = s.clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
			s.Require().NotZero(res.Code, out)
		})
	}
}

func (s *CLITestSuite) TestNewTxCmdUpdateClass() {
	testCases := map[string]struct {
		args []string
		err  error
	}{
		// "valid failing request": {
		// 	args: []string{
		// 		internft.ClassIDFromOwner(s.customer),
		// 	},
		// },
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
			cmd := cli.NewTxCmdUpdateClass()

			out, err := clitestutil.ExecTestCLICmd(s.clientCtx, cmd, append(tc.args, s.commonTxFlags...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res sdk.TxResponse
			err = s.clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
			s.Require().NotZero(res.Code, out)
		})
	}
}

func (s *CLITestSuite) TestNewTxCmdMintNFT() {
	testCases := map[string]struct {
		args []string
		err  error
	}{
		// "valid failing request": {
		// 	args: []string{
		// 		internft.ClassIDFromOwner(s.customer),
		// 		"[]",
		// 		s.customer.String(),
		// 	},
		// },
		"invalid class id": {
			args: []string{
				"",
				"[]",
				s.customer.String(),
			},
			err: internft.ErrInvalidClassID,
		},
		"invalid properties": {
			args: []string{
				internft.ClassIDFromOwner(s.customer),
				"",
				s.customer.String(),
			},
			err: sdkerrors.ErrInvalidType,
		},
		"invalid msg": {
			args: []string{
				internft.ClassIDFromOwner(s.customer),
				"[{}]",
				s.customer.String(),
			},
			err: internft.ErrInvalidTraitID,
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewTxCmdMintNFT()

			out, err := clitestutil.ExecTestCLICmd(s.clientCtx, cmd, append(tc.args, s.commonTxFlags...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res sdk.TxResponse
			err = s.clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
			s.Require().NotZero(res.Code, out)
		})
	}
}

func (s *CLITestSuite) TestNewTxCmdBurnNFT() {
	testCases := map[string]struct {
		args []string
		err  error
	}{
		// "valid failing request": {
		// 	args: []string{
		// 		s.customer.String(),
		// 		internft.NFT{
		// 			ClassId: internft.ClassIDFromOwner(s.customer),
		// 			Id:      math.OneUint(),
		// 		}.String(),
		// 	},
		// },
		"invalid id": {
			args: []string{
				s.customer.String(),
				"",
			},
			err: sdkerrors.ErrInvalidType,
		},
		"invalid msg": {
			args: []string{
				"",
				internft.NFT{
					ClassId: internft.ClassIDFromOwner(s.customer),
					Id:      math.OneUint(),
				}.String(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewTxCmdBurnNFT()

			out, err := clitestutil.ExecTestCLICmd(s.clientCtx, cmd, append(tc.args, s.commonTxFlags...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res sdk.TxResponse
			err = s.clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
			s.Require().NotZero(res.Code, out)
		})
	}
}

func (s *CLITestSuite) TestNewTxCmdUpdateNFT() {
	properties := `[{"id":"uri"}]`

	testCases := map[string]struct {
		args []string
		err  error
	}{
		// "valid failing request": {
		// 	args: []string{
		// 		internft.NFT{
		// 			ClassId: internft.ClassIDFromOwner(s.customer),
		// 			Id:      math.OneUint(),
		// 		}.String(),
		// 		properties,
		// 	},
		// },
		"invalid id": {
			args: []string{
				"",
				properties,
			},
			err: sdkerrors.ErrInvalidType,
		},
		"invalid class id": {
			args: []string{
				internft.NFT{
					Id: math.OneUint(),
				}.String(),
				properties,
			},
			err: internft.ErrInvalidClassID,
		},
		"invalid properties": {
			args: []string{
				internft.NFT{
					ClassId: internft.ClassIDFromOwner(s.customer),
					Id:      math.OneUint(),
				}.String(),
				"",
			},
			err: sdkerrors.ErrInvalidType,
		},
		"invalid msg": {
			args: []string{
				internft.NFT{
					ClassId: internft.ClassIDFromOwner(s.customer),
					Id:      math.OneUint(),
				}.String(),
				"[{}]",
			},
			err: internft.ErrInvalidTraitID,
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewTxCmdUpdateNFT()

			out, err := clitestutil.ExecTestCLICmd(s.clientCtx, cmd, append(tc.args, s.commonTxFlags...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res sdk.TxResponse
			err = s.clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
			s.Require().NotZero(res.Code, out)
		})
	}
}
