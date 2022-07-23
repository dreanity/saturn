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
		return nil, types.ErrUnprovenRandomnessNotExists
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

	reward := sdk.NewCoin("uhydrogen", sdk.NewInt(types.RandomnessVerificationReward))
	rewards := sdk.NewCoins(reward)
	awardedAddress, _ := sdk.AccAddressFromBech32(msg.Creator)

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, rewards); err != nil {
		return nil, err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, awardedAddress, rewards); err != nil {
		return nil, err
	}

	return &types.MsgProveRandomnessResponse{
		Proven: true,
	}, nil
}
