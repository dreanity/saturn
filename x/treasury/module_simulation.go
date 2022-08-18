package treasury

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/dreanity/saturn/testutil/sample"
	treasurysimulation "github.com/dreanity/saturn/x/treasury/simulation"
	"github.com/dreanity/saturn/x/treasury/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = treasurysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgChangeGasPrices = "op_weight_msg_change_gas_prices"
	// TODO: Determine the simulation weight value
	defaultWeightMsgChangeGasPrices int = 100

	opWeightMsgExecuteGasBid = "op_weight_msg_execute_gas_bid"
	// TODO: Determine the simulation weight value
	defaultWeightMsgExecuteGasBid int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	treasuryGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&treasuryGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgChangeGasPrices int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgChangeGasPrices, &weightMsgChangeGasPrices, nil,
		func(_ *rand.Rand) {
			weightMsgChangeGasPrices = defaultWeightMsgChangeGasPrices
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgChangeGasPrices,
		treasurysimulation.SimulateMsgChangeGasPrices(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgExecuteGasBid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgExecuteGasBid, &weightMsgExecuteGasBid, nil,
		func(_ *rand.Rand) {
			weightMsgExecuteGasBid = defaultWeightMsgExecuteGasBid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgExecuteGasBid,
		treasurysimulation.SimulateMsgExecuteGasBid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
