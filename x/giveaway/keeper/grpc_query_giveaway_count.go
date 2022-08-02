package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/giveaway/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GiveawayCount(c context.Context, req *types.QueryGetGiveawayCountRequest) (*types.QueryGetGiveawayCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val := k.GetGiveawayCount(ctx)

	return &types.QueryGetGiveawayCountResponse{GiveawayCount: val}, nil
}
