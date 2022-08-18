package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/dreanity/saturn/testutil/keeper"
	"github.com/dreanity/saturn/testutil/nullify"
	"github.com/dreanity/saturn/x/treasury/keeper"
	"github.com/dreanity/saturn/x/treasury/types"
)

func createTestTreasurer(keeper *keeper.Keeper, ctx sdk.Context) types.Treasurer {
	item := types.Treasurer{}
	keeper.SetTreasurer(ctx, item)
	return item
}

func TestTreasurerGet(t *testing.T) {
	keeper, ctx := keepertest.TreasuryKeeper(t)
	item := createTestTreasurer(keeper, ctx)
	rst, found := keeper.GetTreasurer(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestTreasurerRemove(t *testing.T) {
	keeper, ctx := keepertest.TreasuryKeeper(t)
	createTestTreasurer(keeper, ctx)
	keeper.RemoveTreasurer(ctx)
	_, found := keeper.GetTreasurer(ctx)
	require.False(t, found)
}
