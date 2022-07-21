package randomness_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "saturn/testutil/keeper"
	"saturn/testutil/nullify"
	"saturn/x/randomness"
	"saturn/x/randomness/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RandomnessKeeper(t)
	randomness.InitGenesis(ctx, *k, genesisState)
	got := randomness.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
