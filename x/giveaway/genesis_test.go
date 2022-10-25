package giveaway_test

import (
	"testing"

	keepertest "github.com/dreanity/saturn/testutil/keeper"
	"github.com/dreanity/saturn/testutil/nullify"
	"github.com/dreanity/saturn/x/giveaway"
	"github.com/dreanity/saturn/x/giveaway/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		GiveawayList: []types.Giveaway{
			{
				Index: 0,
			},
			{
				Index: 1,
			},
		},
		GiveawayCount: types.GiveawayCount{
			Value: 74,
		},
		GiveawayByHeightList: []types.GiveawayByHeight{
			{
				Height: 0,
			},
			{
				Height: 1,
			},
		},
		GiveawayByRandomnessList: []types.GiveawayByRandomness{
			{
				Round: 0,
			},
			{
				Round: 1,
			},
		},
		TicketList: []types.Ticket{
			{
				Index: 0,
			},
			{
				Index: 1,
			},
		},
		TicketCountList: []types.TicketCount{
			{
				GiveawayId: 0,
			},
			{
				GiveawayId: 1,
			},
		},
		GiveawaysCountByOrganizerList: []types.GiveawaysCountByOrganizer{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GiveawayKeeper(t)
	giveaway.InitGenesis(ctx, *k, genesisState)
	got := giveaway.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.GiveawayList, got.GiveawayList)
	require.Equal(t, genesisState.GiveawayCount, got.GiveawayCount)
	require.ElementsMatch(t, genesisState.GiveawayByHeightList, got.GiveawayByHeightList)
	require.ElementsMatch(t, genesisState.GiveawayByRandomnessList, got.GiveawayByRandomnessList)
	require.ElementsMatch(t, genesisState.TicketList, got.TicketList)
	require.ElementsMatch(t, genesisState.TicketCountList, got.TicketCountList)
	require.ElementsMatch(t, genesisState.GiveawaysCountByOrganizerList, got.GiveawaysCountByOrganizerList)
	// this line is used by starport scaffolding # genesis/test/assert
}
