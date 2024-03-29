package randomness_test

import (
	"testing"

	keepertest "github.com/dreanity/saturn/testutil/keeper"
	"github.com/dreanity/saturn/testutil/nullify"
	"github.com/dreanity/saturn/x/randomness"
	"github.com/dreanity/saturn/x/randomness/types"

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
				Round: 0,
			},
			{
				Round: 1,
			},
		},
		ProvenRandomnessList: []types.ProvenRandomness{
			{
				Round: 0,
			},
			{
				Round: 1,
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
