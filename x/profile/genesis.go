package profile

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/profile/keeper"
	"github.com/dreanity/saturn/x/profile/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the profile
	for _, elem := range genState.ProfileList {
		k.SetProfile(ctx, elem)
	}
	// Set all the nameRegistry
	for _, elem := range genState.NameRegistryList {
		k.SetNameRegistry(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ProfileList = k.GetAllProfile(ctx)
	genesis.NameRegistryList = k.GetAllNameRegistry(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
