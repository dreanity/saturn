package treasury

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/treasury/keeper"
	"github.com/dreanity/saturn/x/treasury/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	k.SetTreasurer(ctx, genState.Treasurer)
	// Set all the gasPrice
	for _, elem := range genState.GasPriceList {
		k.SetGasPrice(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all treasurer
	treasurer := k.GetTreasurer(ctx)
	genesis.Treasurer = treasurer
	genesis.GasPriceList = k.GetAllGasPrice(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
