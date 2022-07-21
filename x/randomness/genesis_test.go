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

		ChainInfo: &types.ChainInfo{
			PublicKey:   "99",
			Period:      49,
			GenesisTime: 3,
			Hash:        "91",
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RandomnessKeeper(t)
	randomness.InitGenesis(ctx, *k, genesisState)
	got := randomness.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.ChainInfo, got.ChainInfo)
	// this line is used by starport scaffolding # genesis/test/assert
}
