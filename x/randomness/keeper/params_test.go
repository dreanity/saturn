package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "saturn/testutil/keeper"
	"saturn/x/randomness/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RandomnessKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
