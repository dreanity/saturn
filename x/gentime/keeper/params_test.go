package keeper_test

import (
	"testing"

	testkeeper "github.com/dreanity/saturn/testutil/keeper"
	"github.com/dreanity/saturn/x/gentime/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.GentimeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
