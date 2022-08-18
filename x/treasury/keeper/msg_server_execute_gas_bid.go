package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/treasury/types"
)

func (k msgServer) ExecuteGasBid(goCtx context.Context, msg *types.MsgExecuteGasBid) (*types.MsgExecuteGasBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgExecuteGasBidResponse{}, nil
}
