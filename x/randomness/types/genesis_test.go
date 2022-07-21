package types_test

import (
	"testing"

	"saturn/x/randomness/types"

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
				UnprovenRendomnessList: []types.UnprovenRendomness{
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated unprovenRendomness",
			genState: &types.GenesisState{
				UnprovenRendomnessList: []types.UnprovenRendomness{
					{
						Index: "0",
					},
					{
						Index: "0",
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
						Index: "0",
					},
					{
						Index: "0",
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
