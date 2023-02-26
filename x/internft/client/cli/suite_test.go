package cli_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/suite"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtcli "github.com/cometbft/cometbft/libs/cli"
	rpcclientmock "github.com/cometbft/cometbft/rpc/client/mock"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	testutilmod "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/internft"
	"github.com/cosmos/cosmos-sdk/x/internft/client/cli"
	"github.com/cosmos/cosmos-sdk/x/internft/module"
)

type CLITestSuite struct {
	suite.Suite

	encCfg    testutilmod.TestEncodingConfig
	kr        keyring.Keyring
	baseCtx   client.Context
	clientCtx client.Context

	commonTxFlags    []string
	commonQueryFlags []string

	// authority sdk.AccAddress

	vendor   sdk.AccAddress
	customer sdk.AccAddress

	numNFTs uint64
}

func TestCLITestSuite(t *testing.T) {
	suite.Run(t, new(CLITestSuite))
}

func (s *CLITestSuite) SetupSuite() {
	s.T().Log("setting up cli test suite")

	s.encCfg = testutilmod.MakeTestEncodingConfig(module.AppModuleBasic{})
	s.kr = keyring.NewInMemory(s.encCfg.Codec)
	s.baseCtx = client.Context{}.
		WithKeyring(s.kr).
		WithTxConfig(s.encCfg.TxConfig).
		WithCodec(s.encCfg.Codec).
		WithClient(clitestutil.MockCometRPC{Client: rpcclientmock.Client{}}).
		WithAccountRetriever(client.MockAccountRetriever{}).
		WithOutput(io.Discard).
		WithChainID("test-chain")

	s.commonTxFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100)))),
	}
	s.commonQueryFlags = []string{
		fmt.Sprintf("--%s=json", cmtcli.OutputFlag),
	}

	var outBuf bytes.Buffer
	ctxGen := func() client.Context {
		bz, _ := s.encCfg.Codec.Marshal(&sdk.TxResponse{})
		c := clitestutil.NewMockCometRPC(abci.ResponseQuery{
			Value: bz,
		})
		return s.baseCtx.WithClient(c)
	}
	s.clientCtx = ctxGen().WithOutput(&outBuf)

	// create accounts
	addresses := []*sdk.AccAddress{
		&s.vendor,
		&s.customer,
	}
	for i, account := range testutil.CreateKeyringAccounts(s.T(), s.clientCtx.Keyring, len(addresses)) {
		*addresses[i] = account.Address
	}

	// vendor creates a class
	classID := internft.ClassIDFromOwner(s.vendor)
	s.newClass(s.vendor)

	// vendor mints nfts to all accounts by amount of numNFTs
	s.numNFTs = 2

	for _, owner := range []sdk.AccAddress{
		s.vendor,
		s.customer,
	} {
		for range make([]struct{}, s.numNFTs) {
			s.mintNFT(classID, owner)
		}
	}
}

func (s *CLITestSuite) newClass(owner sdk.AccAddress) {
	args := append([]string{
		owner.String(),
		"[]",
	}, s.commonTxFlags...)

	out, err := clitestutil.ExecTestCLICmd(s.clientCtx, cli.NewTxCmdNewClass(), args)
	s.Assert().NoError(err)

	var res sdk.TxResponse
	err = s.clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
	s.Assert().NoError(err, out.String())
	s.Assert().Zero(res.Code, out.String())
}

func (s *CLITestSuite) mintNFT(classID string, recipient sdk.AccAddress) {
	owner := internft.ClassOwner(classID)
	args := append([]string{
		owner.String(),
		"[]",
		recipient.String(),
	}, s.commonTxFlags...)

	out, err := clitestutil.ExecTestCLICmd(s.clientCtx, cli.NewTxCmdMintNFT(), args)
	s.Assert().NoError(err)

	var res sdk.TxResponse
	err = s.clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
	s.Assert().NoError(err, out.String())
	s.Assert().Zero(res.Code, out.String())
}
