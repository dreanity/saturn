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

func createNGiveawayByRandomness(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.GiveawayByRandomness {
	items := make([]types.GiveawayByRandomness, n)
	for i := range items {
		items[i].Round = uint64(i)

		keeper.SetGiveawayByRandomness(ctx, items[i])
	}
	return items
}

func TestGiveawayByRandomnessGet(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	items := createNGiveawayByRandomness(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGiveawayByRandomness(ctx,
			item.Round,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestGiveawayByRandomnessRemove(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	items := createNGiveawayByRandomness(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGiveawayByRandomness(ctx,
			item.Round,
		)
		_, found := keeper.GetGiveawayByRandomness(ctx,
			item.Round,
		)
		require.False(t, found)
	}
}

func TestGiveawayByRandomnessGetAll(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	items := createNGiveawayByRandomness(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGiveawayByRandomness(ctx)),
	)
}
