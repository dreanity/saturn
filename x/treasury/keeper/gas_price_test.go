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

func createNGasPrice(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.GasPrice {
	items := make([]types.GasPrice, n)
	for i := range items {
		items[i].Chain = strconv.Itoa(i)
		items[i].TokenAddress = strconv.Itoa(i)
		items[i].TokenSymbol = "UNKNOWN"
		items[i].Value = "1"

		keeper.SetGasPrice(ctx, items[i])
	}
	return items
}

func TestGasPriceGet(t *testing.T) {
	keeper, ctx := keepertest.TreasuryKeeper(t)
	items := createNGasPrice(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGasPrice(ctx,
			item.Chain,
			item.TokenAddress,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestGasPriceRemove(t *testing.T) {
	keeper, ctx := keepertest.TreasuryKeeper(t)
	items := createNGasPrice(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGasPrice(ctx,
			item.Chain,
			item.TokenAddress,
		)
		_, found := keeper.GetGasPrice(ctx,
			item.Chain,
			item.TokenAddress,
		)
		require.False(t, found)
	}
}

func TestGasPriceGetAll(t *testing.T) {
	keeper, ctx := keepertest.TreasuryKeeper(t)
	items := createNGasPrice(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGasPrice(ctx)),
	)
}
