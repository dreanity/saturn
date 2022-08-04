package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/dreanity/saturn/x/giveaway/keeper"
	"github.com/dreanity/saturn/x/giveaway/types"
)

func SimulateMsgIssueTicket(
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgIssueTicket{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the IssueTicket simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "IssueTicket simulation not implemented"), nil, nil
	}
}
