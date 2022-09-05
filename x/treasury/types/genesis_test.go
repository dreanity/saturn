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
						Chain:        "0",
						TokenAddress: "0",
						TokenSymbol:  "UNKNOWN",
						Value:        "1",
					},
					{
						Chain:        "1",
						TokenAddress: "0",
						TokenSymbol:  "UNKNOWN",
						Value:        "1",
					},
				},
				GasBidList: []types.GasBid{
					{
						Chain: "0",
					},
					{
						Chain: "1",
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
						Chain:        "0",
						TokenAddress: "0",
						TokenSymbol:  "UNKNOWN",
						Value:        "1",
					},
					{
						Chain:        "0",
						TokenAddress: "0",
						TokenSymbol:  "UNKNOWN",
						Value:        "1",
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
						Chain: "0",
					},
					{
						Chain: "0",
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
