package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/dreanity/saturn/testutil/keeper"
	"github.com/dreanity/saturn/testutil/nullify"
	"github.com/dreanity/saturn/x/giveaway/keeper"
	"github.com/dreanity/saturn/x/giveaway/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTicketCount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TicketCount {
	items := make([]types.TicketCount, n)
	for i := range items {
		items[i].GiveawayId = uint32(i)

		keeper.SetTicketCount(ctx, items[i])
	}
	return items
}

func TestTicketCountGet(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	items := createNTicketCount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTicketCount(ctx,
			item.GiveawayId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTicketCountRemove(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	items := createNTicketCount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTicketCount(ctx,
			item.GiveawayId,
		)
		_, found := keeper.GetTicketCount(ctx,
			item.GiveawayId,
		)
		require.False(t, found)
	}
}

func TestTicketCountGetAll(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	items := createNTicketCount(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTicketCount(ctx)),
	)
}
