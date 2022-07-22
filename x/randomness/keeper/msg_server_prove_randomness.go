package keeper

import (
	"context"

	"saturn/x/randomness/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ProveRandomness(goCtx context.Context, msg *types.MsgProveRandomness) (*types.MsgProveRandomnessResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgProveRandomnessResponse{}, nil
}
