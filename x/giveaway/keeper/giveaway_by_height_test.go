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

func createNGiveawayByHeight(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.GiveawayByHeight {
	items := make([]types.GiveawayByHeight, n)
	for i := range items {
		items[i].Height = int32(i)

		keeper.SetGiveawayByHeight(ctx, items[i])
	}
	return items
}

func TestGiveawayByHeightGet(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	items := createNGiveawayByHeight(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGiveawayByHeight(ctx,
			item.Height,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestGiveawayByHeightRemove(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	items := createNGiveawayByHeight(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGiveawayByHeight(ctx,
			item.Height,
		)
		_, found := keeper.GetGiveawayByHeight(ctx,
			item.Height,
		)
		require.False(t, found)
	}
}

func TestGiveawayByHeightGetAll(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	items := createNGiveawayByHeight(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGiveawayByHeight(ctx)),
	)
}
