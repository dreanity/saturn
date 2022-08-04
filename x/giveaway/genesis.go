package giveaway

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/giveaway/keeper"
	"github.com/dreanity/saturn/x/giveaway/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the giveaway
	for _, elem := range genState.GiveawayList {
		k.SetGiveaway(ctx, elem)
	}
	// Set if defined
	k.SetGiveawayCount(ctx, genState.GiveawayCount)
	// Set all the giveawayByHeight
	for _, elem := range genState.GiveawayByHeightList {
		k.SetGiveawayByHeight(ctx, elem)
	}
	// Set all the giveawayByRandomness
	for _, elem := range genState.GiveawayByRandomnessList {
		k.SetGiveawayByRandomness(ctx, elem)
	}
	// Set all the ticket
	for _, elem := range genState.TicketList {
		k.SetTicket(ctx, elem)
	}
	// Set all the ticketCount
	for _, elem := range genState.TicketCountList {
		k.SetTicketCount(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.GiveawayList = k.GetAllGiveaway(ctx)
	// Get all giveawayCount
	genesis.GiveawayCount = k.GetGiveawayCount(ctx)
	genesis.GiveawayByHeightList = k.GetAllGiveawayByHeight(ctx)
	genesis.GiveawayByRandomnessList = k.GetAllGiveawayByRandomness(ctx)
	genesis.TicketList = k.GetAllTicket(ctx)
	genesis.TicketCountList = k.GetAllTicketCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
