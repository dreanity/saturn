package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/treasury/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Treasurer(c context.Context, req *types.QueryGetTreasurerRequest) (*types.QueryGetTreasurerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val := k.GetTreasurer(ctx)

	return &types.QueryGetTreasurerResponse{Treasurer: val}, nil
}
