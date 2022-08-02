package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/dreanity/saturn/testutil/keeper"
	"github.com/dreanity/saturn/testutil/nullify"
	"github.com/dreanity/saturn/x/gentime/keeper"
	"github.com/dreanity/saturn/x/gentime/types"
)

func createTestTime(keeper *keeper.Keeper, ctx sdk.Context) types.Time {
	item := types.Time{}
	keeper.SetTime(ctx, item)
	return item
}

func TestTimeGet(t *testing.T) {
	keeper, ctx := keepertest.GentimeKeeper(t)
	item := createTestTime(keeper, ctx)
	rst, found := keeper.GetTime(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestTimeRemove(t *testing.T) {
	keeper, ctx := keepertest.GentimeKeeper(t)
	createTestTime(keeper, ctx)
	keeper.RemoveTime(ctx)
	_, found := keeper.GetTime(ctx)
	require.False(t, found)
}
