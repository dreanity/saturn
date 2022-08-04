package giveaway

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/dreanity/saturn/testutil/sample"
	giveawaysimulation "github.com/dreanity/saturn/x/giveaway/simulation"
	"github.com/dreanity/saturn/x/giveaway/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = giveawaysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateGiveaway = "op_weight_msg_create_giveaway"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateGiveaway int = 100

	opWeightMsgIssueTicket = "op_weight_msg_issue_ticket"
	// TODO: Determine the simulation weight value
	defaultWeightMsgIssueTicket int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	giveawayGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&giveawayGenesis)
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

	var weightMsgCreateGiveaway int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateGiveaway, &weightMsgCreateGiveaway, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGiveaway = defaultWeightMsgCreateGiveaway
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateGiveaway,
		giveawaysimulation.SimulateMsgCreateGiveaway(am.gentimeKeeper, am.keeper),
	))

	var weightMsgIssueTicket int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgIssueTicket, &weightMsgIssueTicket, nil,
		func(_ *rand.Rand) {
			weightMsgIssueTicket = defaultWeightMsgIssueTicket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgIssueTicket,
		giveawaysimulation.SimulateMsgIssueTicket(am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
