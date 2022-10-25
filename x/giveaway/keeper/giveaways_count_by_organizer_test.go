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

func createNGiveawaysCountByOrganizer(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.GiveawaysCountByOrganizer {
	items := make([]types.GiveawaysCountByOrganizer, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetGiveawaysCountByOrganizer(ctx, items[i])
	}
	return items
}

func TestGiveawaysCountByOrganizerGet(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	items := createNGiveawaysCountByOrganizer(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGiveawaysCountByOrganizer(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestGiveawaysCountByOrganizerRemove(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	items := createNGiveawaysCountByOrganizer(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGiveawaysCountByOrganizer(ctx,
			item.Address,
		)
		_, found := keeper.GetGiveawaysCountByOrganizer(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestGiveawaysCountByOrganizerGetAll(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	items := createNGiveawaysCountByOrganizer(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGiveawaysCountByOrganizer(ctx)),
	)
}
