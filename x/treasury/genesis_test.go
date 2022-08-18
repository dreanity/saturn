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
				Currency: "0",
			},
			{
				Currency: "1",
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
	// this line is used by starport scaffolding # genesis/test/assert
}
