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

func createNProvenRandomness(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ProvenRandomness {
	items := make([]types.ProvenRandomness, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetProvenRandomness(ctx, items[i])
	}
	return items
}

func TestProvenRandomnessGet(t *testing.T) {
	keeper, ctx := keepertest.RandomnessKeeper(t)
	items := createNProvenRandomness(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProvenRandomness(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestProvenRandomnessRemove(t *testing.T) {
	keeper, ctx := keepertest.RandomnessKeeper(t)
	items := createNProvenRandomness(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProvenRandomness(ctx,
			item.Index,
		)
		_, found := keeper.GetProvenRandomness(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestProvenRandomnessGetAll(t *testing.T) {
	keeper, ctx := keepertest.RandomnessKeeper(t)
	items := createNProvenRandomness(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProvenRandomness(ctx)),
	)
}
