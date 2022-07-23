package keeper

import (
	"context"

	"saturn/x/randomness/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ProveRandomness(goCtx context.Context, msg *types.MsgProveRandomness) (*types.MsgProveRandomnessResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	gasMeter := ctx.GasMeter()
	gasMeter.ConsumeGas(types.RandomnessVerificationGas, "randomness verification")

	unprovenRandomness, isFound := k.GetUnprovenRandomness(ctx, msg.Round)
	if !isFound {
		unprovenRandomness = types.UnprovenRandomness{Round: msg.Round}
		k.SetUnprovenRandomness(ctx, unprovenRandomness)

		// return nil, types.ErrUnprovenRandomnessNotExists
	}

	provenRandomness := types.ProvenRandomness{
		Round:             msg.Round,
		Randomness:        msg.Randomness,
		Signature:         msg.Signature,
		PreviousSignature: msg.PreviousSignature,
	}

	chainInfo := k.GetChainInfo(ctx)

	if err := provenRandomness.Verify(&chainInfo); err != nil {
		return nil, types.ErrRandomnessVerification.Wrap(err.Error())
	}

	k.RemoveUnprovenRandomness(ctx, unprovenRandomness.Round)
	k.SetProvenRandomness(ctx, provenRandomness)

	randomnessProved := types.RandomnessProved{
		Round: provenRandomness.Round,
	}

	ctx.EventManager().EmitTypedEvent(&randomnessProved)

	gasMeter.RefundGas(gasMeter.GasConsumed(), "Refund for proven randomness")

	return &types.MsgProveRandomnessResponse{
		Proven: true,
	}, nil
}
