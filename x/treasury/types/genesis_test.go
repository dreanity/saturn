package types_test

import (
	"testing"

	"github.com/dreanity/saturn/x/treasury/types"
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

				Treasurer: types.Treasurer{
					Address: "56",
				},
				GasPriceList: []types.GasPrice{
					{
						Currency: "0",
					},
					{
						Currency: "1",
					},
				},
				GasBidList: []types.GasBid{
					{
						FromChain: "0",
					},
					{
						FromChain: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated gasPrice",
			genState: &types.GenesisState{
				GasPriceList: []types.GasPrice{
					{
						Currency: "0",
					},
					{
						Currency: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated gasBid",
			genState: &types.GenesisState{
				GasBidList: []types.GasBid{
					{
						FromChain: "0",
					},
					{
						FromChain: "0",
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
