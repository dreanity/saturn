package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/dreanity/saturn/testutil/keeper"
	"github.com/dreanity/saturn/testutil/nullify"
	"github.com/dreanity/saturn/x/treasury/keeper"
	"github.com/dreanity/saturn/x/treasury/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNGasBid(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.GasBid {
	items := make([]types.GasBid, n)
	for i := range items {
		items[i].FromChain = strconv.Itoa(i)

		keeper.SetGasBid(ctx, items[i])
	}
	return items
}

func TestGasBidGet(t *testing.T) {
	keeper, ctx := keepertest.TreasuryKeeper(t)
	items := createNGasBid(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGasBid(ctx,
			item.FromChain,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestGasBidRemove(t *testing.T) {
	keeper, ctx := keepertest.TreasuryKeeper(t)
	items := createNGasBid(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGasBid(ctx,
			item.FromChain,
		)
		_, found := keeper.GetGasBid(ctx,
			item.FromChain,
		)
		require.False(t, found)
	}
}

func TestGasBidGetAll(t *testing.T) {
	keeper, ctx := keepertest.TreasuryKeeper(t)
	items := createNGasBid(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGasBid(ctx)),
	)
}
