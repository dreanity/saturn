package treasury_test

import (
	"testing"

	keepertest "github.com/dreanity/saturn/testutil/keeper"
	"github.com/dreanity/saturn/testutil/nullify"
	"github.com/dreanity/saturn/x/treasury"
	"github.com/dreanity/saturn/x/treasury/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		Treasurer: types.Treasurer{
			Address: "96",
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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TreasuryKeeper(t)
	treasury.InitGenesis(ctx, *k, genesisState)
	got := treasury.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.Treasurer, got.Treasurer)
	require.ElementsMatch(t, genesisState.GasPriceList, got.GasPriceList)
	require.ElementsMatch(t, genesisState.GasBidList, got.GasBidList)
	// this line is used by starport scaffolding # genesis/test/assert
}
