package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/dreanity/saturn/testutil/keeper"
	"github.com/dreanity/saturn/testutil/nullify"
	"github.com/dreanity/saturn/x/giveaway/keeper"
	"github.com/dreanity/saturn/x/giveaway/types"
)

func createTestGiveawayCount(keeper *keeper.Keeper, ctx sdk.Context) types.GiveawayCount {
	item := types.GiveawayCount{}
	keeper.SetGiveawayCount(ctx, item)
	return item
}

func TestGiveawayCountGet(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	item := createTestGiveawayCount(keeper, ctx)
	rst := keeper.GetGiveawayCount(ctx)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}
