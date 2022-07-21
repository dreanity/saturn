package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"saturn/x/randomness/types"
)

func (k Keeper) UnprovenRendomnessAll(c context.Context, req *types.QueryAllUnprovenRendomnessRequest) (*types.QueryAllUnprovenRendomnessResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var unprovenRendomnesss []types.UnprovenRendomness
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	unprovenRendomnessStore := prefix.NewStore(store, types.KeyPrefix(types.UnprovenRendomnessKeyPrefix))

	pageRes, err := query.Paginate(unprovenRendomnessStore, req.Pagination, func(key []byte, value []byte) error {
		var unprovenRendomness types.UnprovenRendomness
		if err := k.cdc.Unmarshal(value, &unprovenRendomness); err != nil {
			return err
		}

		unprovenRendomnesss = append(unprovenRendomnesss, unprovenRendomness)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUnprovenRendomnessResponse{UnprovenRendomness: unprovenRendomnesss, Pagination: pageRes}, nil
}

func (k Keeper) UnprovenRendomness(c context.Context, req *types.QueryGetUnprovenRendomnessRequest) (*types.QueryGetUnprovenRendomnessResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetUnprovenRendomness(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetUnprovenRendomnessResponse{UnprovenRendomness: val}, nil
}
