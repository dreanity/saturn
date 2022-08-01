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

func createNGiveaway(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Giveaway {
	items := make([]types.Giveaway, n)
	for i := range items {
		items[i].Index = uint64(i)

		keeper.SetGiveaway(ctx, items[i])
	}
	return items
}

func TestGiveawayGet(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	items := createNGiveaway(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGiveaway(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestGiveawayRemove(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	items := createNGiveaway(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGiveaway(ctx,
			item.Index,
		)
		_, found := keeper.GetGiveaway(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestGiveawayGetAll(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	items := createNGiveaway(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGiveaway(ctx)),
	)
}
