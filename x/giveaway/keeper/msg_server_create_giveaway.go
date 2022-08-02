package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/giveaway/types"
)

func (k msgServer) CreateGiveaway(goCtx context.Context, msg *types.MsgCreateGiveaway) (*types.MsgCreateGiveawayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_ = ctx

	// if msg.Duration >

	// giveway := types.Giveaway{

	// }

	return &types.MsgCreateGiveawayResponse{}, nil
}
