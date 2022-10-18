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

func createNProfile(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Profile {
	items := make([]types.Profile, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetProfile(ctx, items[i])
	}
	return items
}

func TestProfileGet(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	items := createNProfile(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProfile(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestProfileRemove(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	items := createNProfile(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProfile(ctx,
			item.Address,
		)
		_, found := keeper.GetProfile(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestProfileGetAll(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	items := createNProfile(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProfile(ctx)),
	)
}
