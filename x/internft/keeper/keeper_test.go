package keeper_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/suite"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttime "github.com/cometbft/cometbft/types/time"

	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/internft"
	"github.com/cosmos/cosmos-sdk/x/internft/keeper"
	"github.com/cosmos/cosmos-sdk/x/internft/module"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx sdk.Context

	keeper keeper.Keeper

	queryServer internft.QueryServer
	msgServer   internft.MsgServer

	vendor   sdk.AccAddress
	customer sdk.AccAddress

	mutableTraitID   string
	immutableTraitID string

	numNFTs uint64
}

func createAddresses(size int, prefix string) []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, size)
	for i := range addrs {
		addrs[i] = sdk.AccAddress(fmt.Sprintf("%s%d", prefix, i))
	}

	return addrs
}

func createClassIDs(size int, prefix string) []string {
	owners := createAddresses(size, prefix)
	ids := make([]string, len(owners))
	for i, owner := range owners {
		ids[i] = internft.ClassIDFromOwner(owner)
	}

	return ids
}

func randomString(size int) string {
	res := make([]rune, size)

	letters := []rune("0123456789abcdef")
	for i := range res {
		res[i] = letters[rand.Intn(len(letters))]
	}

	return string(res)
}

func (s *KeeperTestSuite) SetupTest() {
	key := storetypes.NewKVStoreKey(internft.StoreKey)
	encCfg := moduletestutil.MakeTestEncodingConfig(module.AppModuleBasic{})
	s.keeper = keeper.NewKeeper(key, encCfg.Codec, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	testCtx := testutil.DefaultContextWithDB(s.T(), key, storetypes.NewTransientStoreKey("transient_test"))
	s.ctx = testCtx.Ctx.WithBlockHeader(cmtproto.Header{Time: cmttime.Now()})

	s.queryServer = keeper.NewQueryServer(s.keeper)
	s.msgServer = keeper.NewMsgServer(s.keeper)

	// create accounts
	addresses := []*sdk.AccAddress{
		&s.vendor,
		&s.customer,
	}
	for i, address := range createAddresses(len(addresses), "addr") {
		*addresses[i] = address
	}

	s.keeper.SetParams(s.ctx, internft.Params{})

	// vendor creates a class
	class := internft.Class{
		Id: internft.ClassIDFromOwner(s.vendor),
	}
	err := class.ValidateBasic()
	s.Assert().NoError(err)

	s.mutableTraitID = "level"
	s.immutableTraitID = "color"

	traits := []internft.Trait{
		{
			Id:      s.mutableTraitID,
			Mutable: true,
		},
		{
			Id: s.immutableTraitID,
		},
	}

	err = s.keeper.NewClass(s.ctx, class, traits)
	s.Assert().NoError(err)

	// vendor mints nfts to all accounts by amount of numNFTs
	s.numNFTs = 2

	for _, owner := range []sdk.AccAddress{
		s.vendor,
		s.customer,
	} {
		for range make([]struct{}, s.numNFTs) {
			properties := []internft.Property{
				{
					Id: s.mutableTraitID,
				},
				{
					Id: s.immutableTraitID,
				},
			}

			id, err := s.keeper.MintNFT(s.ctx, owner, class.Id, properties)
			s.Assert().NoError(err)
			s.Assert().NotNil(id)
		}
	}
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
