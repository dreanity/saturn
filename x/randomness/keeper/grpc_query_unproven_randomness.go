package keeper

import (
	"context"

	"saturn/x/randomness/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UnprovenRandomnessAll(c context.Context, req *types.QueryAllUnprovenRandomnessRequest) (*types.QueryAllUnprovenRandomnessResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var unprovenRandomnesss []types.UnprovenRandomness
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	unprovenRandomnessStore := prefix.NewStore(store, types.KeyPrefix(types.UnprovenRandomnessKeyPrefix))

	pageRes, err := query.Paginate(unprovenRandomnessStore, req.Pagination, func(key []byte, value []byte) error {
		var unprovenRandomness types.UnprovenRandomness
		if err := k.cdc.Unmarshal(value, &unprovenRandomness); err != nil {
			return err
		}

		unprovenRandomnesss = append(unprovenRandomnesss, unprovenRandomness)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUnprovenRandomnessResponse{UnprovenRandomness: unprovenRandomnesss, Pagination: pageRes}, nil
}

func (k Keeper) UnprovenRandomness(c context.Context, req *types.QueryGetUnprovenRandomnessRequest) (*types.QueryGetUnprovenRandomnessResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetUnprovenRandomness(
		ctx,
		req.Round,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetUnprovenRandomnessResponse{UnprovenRandomness: val}, nil
}
