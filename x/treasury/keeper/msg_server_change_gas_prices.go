package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/treasury/types"
)

func (k msgServer) ChangeGasPrices(goCtx context.Context, msg *types.MsgChangeGasPrices) (*types.MsgChangeGasPricesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgChangeGasPricesResponse{}, nil
}
