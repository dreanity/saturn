package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/dreanity/saturn/x/randomness/keeper"
	"github.com/dreanity/saturn/x/randomness/types"
)

func SimulateMsgProveRandomness(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgProveRandomness{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ProveRandomness simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ProveRandomness simulation not implemented"), nil, nil
	}
}
