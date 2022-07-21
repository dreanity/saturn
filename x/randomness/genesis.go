package randomness

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"saturn/x/randomness/keeper"
	"saturn/x/randomness/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.ChainInfo != nil {
		k.SetChainInfo(ctx, *genState.ChainInfo)
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
		genesis.ChainInfo = &chainInfo
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
