package profile_test

import (
	"testing"

	keepertest "github.com/dreanity/saturn/testutil/keeper"
	"github.com/dreanity/saturn/testutil/nullify"
	"github.com/dreanity/saturn/x/profile"
	"github.com/dreanity/saturn/x/profile/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ProfileList: []types.Profile{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		NameRegistryList: []types.NameRegistry{
			{
				Name: "0",
			},
			{
				Name: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ProfileKeeper(t)
	profile.InitGenesis(ctx, *k, genesisState)
	got := profile.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ProfileList, got.ProfileList)
	require.ElementsMatch(t, genesisState.NameRegistryList, got.NameRegistryList)
	// this line is used by starport scaffolding # genesis/test/assert
}
