package types_test

import (
	"testing"

	"github.com/dreanity/saturn/x/profile/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated profile",
			genState: &types.GenesisState{
				ProfileList: []types.Profile{
					{
						Address: "0",
					},
					{
						Address: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated nameRegistry",
			genState: &types.GenesisState{
				NameRegistryList: []types.NameRegistry{
					{
						Name: "0",
					},
					{
						Name: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
