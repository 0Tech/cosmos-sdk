package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttime "github.com/cometbft/cometbft/types/time"

	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/testutil"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/internft"
	"github.com/cosmos/cosmos-sdk/x/internft/keeper"
	"github.com/cosmos/cosmos-sdk/x/internft/module"
)

func TestImportExportGenesis(t *testing.T) {
	key := storetypes.NewKVStoreKey(internft.StoreKey)
	encCfg := moduletestutil.MakeTestEncodingConfig(module.AppModuleBasic{})
	keeper := keeper.NewKeeper(key, encCfg.Codec, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	testCtx := testutil.DefaultContextWithDB(t, key, storetypes.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockHeader(cmtproto.Header{Time: cmttime.Now()})

	classIDs := createClassIDs(2, "class")
	addr := createAddresses(1, "addr")[0]

	testCases := map[string]struct {
		gs *internft.GenesisState
	}{
		"default": {
			gs: internft.DefaultGenesisState(),
		},
		"no compositions": {
			gs: &internft.GenesisState{
				Params: internft.DefaultParams(),
				Classes: []internft.GenesisClass{
					{
						Id:              classIDs[0],
						LastMintedNftId: math.NewUint(2),
						Nfts: []internft.GenesisNFT{
							{
								Id:    math.NewUint(1),
								Owner: addr.String(),
							},
							{
								Id:    math.NewUint(2),
								Owner: addr.String(),
							},
						},
					},
					{
						Id:              classIDs[1],
						LastMintedNftId: math.NewUint(2),
						Nfts: []internft.GenesisNFT{
							{
								Id:    math.NewUint(1),
								Owner: addr.String(),
							},
							{
								Id:    math.NewUint(2),
								Owner: addr.String(),
							},
						},
					},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, _ := ctx.CacheContext()

			err := tc.gs.ValidateBasic()
			assert.NoError(t, err)

			err = keeper.InitGenesis(ctx, tc.gs)
			require.NoError(t, err)

			exported := keeper.ExportGenesis(ctx)
			require.Equal(t, tc.gs, exported)
		})
	}
}
