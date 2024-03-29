package types_test

import (
	"testing"

	"github.com/dreanity/saturn/x/randomness/types"

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

				ChainInfo: types.ChainInfo{
					PublicKey:   "6",
					Period:      25,
					GenesisTime: 42,
					Hash:        "22",
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated unprovenRandomness",
			genState: &types.GenesisState{
				UnprovenRandomnessList: []types.UnprovenRandomness{
					{
						Round: 0,
					},
					{
						Round: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated provenRandomness",
			genState: &types.GenesisState{
				ProvenRandomnessList: []types.ProvenRandomness{
					{
						Round: 0,
					},
					{
						Round: 0,
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
