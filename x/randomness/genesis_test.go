package randomness_test

import (
	"testing"

	keepertest "saturn/testutil/keeper"
	"saturn/testutil/nullify"
	"saturn/x/randomness"
	"saturn/x/randomness/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ChainInfo: types.ChainInfo{
			PublicKey:   "99",
			Period:      49,
			GenesisTime: 3,
			Hash:        "91",
		},
		UnprovenRandomnessList: []types.UnprovenRandomness{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		ProvenRandomnessList: []types.ProvenRandomness{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
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
	require.ElementsMatch(t, genesisState.UnprovenRandomnessList, got.UnprovenRandomnessList)
	require.ElementsMatch(t, genesisState.ProvenRandomnessList, got.ProvenRandomnessList)
	// this line is used by starport scaffolding # genesis/test/assert
}
