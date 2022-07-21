package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "saturn/testutil/keeper"
	"saturn/testutil/nullify"
	"saturn/x/randomness/keeper"
	"saturn/x/randomness/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNUnprovenRendomness(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.UnprovenRendomness {
	items := make([]types.UnprovenRendomness, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetUnprovenRendomness(ctx, items[i])
	}
	return items
}

func TestUnprovenRendomnessGet(t *testing.T) {
	keeper, ctx := keepertest.RandomnessKeeper(t)
	items := createNUnprovenRendomness(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUnprovenRendomness(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestUnprovenRendomnessRemove(t *testing.T) {
	keeper, ctx := keepertest.RandomnessKeeper(t)
	items := createNUnprovenRendomness(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUnprovenRendomness(ctx,
			item.Index,
		)
		_, found := keeper.GetUnprovenRendomness(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestUnprovenRendomnessGetAll(t *testing.T) {
	keeper, ctx := keepertest.RandomnessKeeper(t)
	items := createNUnprovenRendomness(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUnprovenRendomness(ctx)),
	)
}
