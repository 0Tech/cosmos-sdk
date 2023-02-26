package module

import (
	"context"
	"encoding/json"
	"fmt"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	storetypes "cosmossdk.io/store/types"

	modulev1alpha1 "github.com/cosmos/cosmos-sdk/andromeda-api/cosmos/internft/module/v1alpha1"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/internft"
	// "github.com/cosmos/cosmos-sdk/x/internft/client/cli"
	"github.com/cosmos/cosmos-sdk/x/internft/keeper"
)

// AppModuleBasic defines the basic application module used by the module.
type AppModuleBasic struct{}

var _ module.AppModuleBasic = (*AppModuleBasic)(nil)

// Name returns the ModuleName
func (AppModuleBasic) Name() string {
	return internft.ModuleName
}

// RegisterLegacyAminoCodec registers the types on the LegacyAmino codec
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

func (b AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	internft.RegisterInterfaces(registry)
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if err := internft.RegisterQueryHandlerClient(context.Background(), mux, internft.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// GetTxCmd returns the transaction commands for the module
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	// return cli.NewTxCmd()
	return nil
}

// GetQueryCmd returns the cli query commands for the module
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	// return cli.NewQueryCmd()
	return nil
}

// ____________________________________________________________________________

var _ module.HasGenesisBasics = (*AppModuleBasic)(nil)

// DefaultGenesis returns default genesis state as raw bytes for the module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(internft.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var gs internft.GenesisState
	if err := cdc.UnmarshalJSON(bz, &gs); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", internft.ModuleName, err)
	}

	return gs.ValidateBasic()
}

// ____________________________________________________________________________

var _ module.AppModule = (*AppModule)(nil)

// AppModule implements an application module for the module.
type AppModule struct {
	AppModuleBasic

	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
	return AppModule{
		keeper: keeper,
	}
}

// ____________________________________________________________________________

var _ module.HasInvariants = (*AppModule)(nil)

// RegisterInvariants does nothing, there are no invariants to enforce
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

// ____________________________________________________________________________

var _ module.HasServices = (*AppModule)(nil)

// RegisterServices registers a GRPC query service to respond to the
// module-specific GRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	internft.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServer(am.keeper))
	internft.RegisterQueryServer(cfg.QueryServer(), keeper.NewQueryServer(am.keeper))

	// m := keeper.NewMigrator(am.keeper)
	// migrations := map[uint64]func(sdk.Context) error{}
	// for ver, handler := range migrations {
	// 	if err := cfg.RegisterMigration(internft.ModuleName, ver, handler); err != nil {
	// 		panic(fmt.Sprintf("failed to migrate x/%s from version %d to %d: %v", internft.ModuleName, ver, ver+1, err))
	// 	}
	// }
}

// ____________________________________________________________________________

var _ module.HasGenesis = (*AppModule)(nil)

// InitGenesis performs genesis initialization for the module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var gs internft.GenesisState
	cdc.MustUnmarshalJSON(data, &gs)

	if err := am.keeper.InitGenesis(ctx, &gs); err != nil {
		panic(err)
	}

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := am.keeper.ExportGenesis(ctx)
	return cdc.MustMarshalJSON(gs)
}

// ____________________________________________________________________________

var _ module.HasConsensusVersion = (*AppModule)(nil)

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 1 }

// ____________________________________________________________________________

var _ appmodule.AppModule = (*AppModule)(nil)

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

//
// App Wiring Setup
//

func init() {
	appmodule.Register(
		&modulev1alpha1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type InterNFTInputs struct {
	depinject.In

	Config *modulev1alpha1.Module
	Key    *storetypes.KVStoreKey
	Cdc    codec.Codec
}

type InterNFTOutputs struct {
	depinject.Out

	Keeper keeper.Keeper
	Module appmodule.AppModule
}

func ProvideModule(in InterNFTInputs) InterNFTOutputs {
	// default to governance authority if not provided
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	k := keeper.NewKeeper(in.Key, in.Cdc, authority.String())
	m := NewAppModule(in.Cdc, k)

	return InterNFTOutputs{Keeper: k, Module: m}
}
