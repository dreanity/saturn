package keeper

import (
	"context"

	"github.com/dreanity/saturn/x/randomness/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ChainInfo(c context.Context, req *types.QueryGetChainInfoRequest) (*types.QueryGetChainInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	val := k.GetChainInfo(ctx)

	return &types.QueryGetChainInfoResponse{ChainInfo: val}, nil
}
