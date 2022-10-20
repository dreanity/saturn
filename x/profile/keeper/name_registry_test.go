package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/dreanity/saturn/testutil/keeper"
	"github.com/dreanity/saturn/testutil/nullify"
	"github.com/dreanity/saturn/x/profile/keeper"
	"github.com/dreanity/saturn/x/profile/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNameRegistry(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NameRegistry {
	items := make([]types.NameRegistry, n)
	for i := range items {
		items[i].Name = strconv.Itoa(i)

		keeper.SetNameRegistry(ctx, items[i])
	}
	return items
}

func TestNameRegistryGet(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	items := createNNameRegistry(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNameRegistry(ctx,
			item.Name,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNameRegistryRemove(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	items := createNNameRegistry(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNameRegistry(ctx,
			item.Name,
		)
		_, found := keeper.GetNameRegistry(ctx,
			item.Name,
		)
		require.False(t, found)
	}
}

func TestNameRegistryGetAll(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	items := createNNameRegistry(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNameRegistry(ctx)),
	)
}
