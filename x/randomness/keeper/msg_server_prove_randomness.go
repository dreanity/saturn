package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"saturn/x/randomness/types"
)

func (k msgServer) ProveRandomness(goCtx context.Context, msg *types.MsgProveRandomness) (*types.MsgProveRandomnessResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgProveRandomnessResponse{}, nil
}
