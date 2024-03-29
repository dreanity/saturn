package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/dreanity/saturn/testutil/keeper"
	"github.com/dreanity/saturn/testutil/nullify"
	"github.com/dreanity/saturn/x/randomness/keeper"
	"github.com/dreanity/saturn/x/randomness/types"
)

func createTestChainInfo(keeper *keeper.Keeper, ctx sdk.Context) types.ChainInfo {
	item := types.ChainInfo{}
	keeper.SetChainInfo(ctx, item)
	return item
}

func TestChainInfoGet(t *testing.T) {
	keeper, ctx := keepertest.RandomnessKeeper(t)
	item := createTestChainInfo(keeper, ctx)
	rst := keeper.GetChainInfo(ctx)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}
