package randomness

import (
	"saturn/x/randomness/keeper"
	"saturn/x/randomness/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	k.SetChainInfo(ctx, genState.ChainInfo)
	// Set all the unprovenRendomness
	for _, elem := range genState.UnprovenRendomnessList {
		k.SetUnprovenRendomness(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all chainInfo
	chainInfo, found := k.GetChainInfo(ctx)
	if found {
		genesis.ChainInfo = chainInfo
	}
	genesis.UnprovenRendomnessList = k.GetAllUnprovenRendomness(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
