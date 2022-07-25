package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/dreanity/saturn/testutil/keeper"
	"github.com/dreanity/saturn/testutil/nullify"
	"github.com/dreanity/saturn/x/randomness/keeper"
	"github.com/dreanity/saturn/x/randomness/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNUnprovenRandomness(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.UnprovenRandomness {
	items := make([]types.UnprovenRandomness, n)
	for i := range items {
		items[i].Round = uint64(i)

		keeper.SetUnprovenRandomness(ctx, items[i])
	}
	return items
}

func TestUnprovenRandomnessGet(t *testing.T) {
	keeper, ctx := keepertest.RandomnessKeeper(t)
	items := createNUnprovenRandomness(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUnprovenRandomness(ctx,
			item.Round,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestUnprovenRandomnessRemove(t *testing.T) {
	keeper, ctx := keepertest.RandomnessKeeper(t)
	items := createNUnprovenRandomness(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUnprovenRandomness(ctx,
			item.Round,
		)
		_, found := keeper.GetUnprovenRandomness(ctx,
			item.Round,
		)
		require.False(t, found)
	}
}

func TestUnprovenRandomnessGetAll(t *testing.T) {
	keeper, ctx := keepertest.RandomnessKeeper(t)
	items := createNUnprovenRandomness(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUnprovenRandomness(ctx)),
	)
}
