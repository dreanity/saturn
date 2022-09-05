package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/dreanity/saturn/x/treasury/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GasBidAll(c context.Context, req *types.QueryAllGasBidRequest) (*types.QueryAllGasBidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var gasBids []types.GasBid
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	gasBidStore := prefix.NewStore(store, types.KeyPrefix(types.GasBidKeyPrefix))

	pageRes, err := query.Paginate(gasBidStore, req.Pagination, func(key []byte, value []byte) error {
		var gasBid types.GasBid
		if err := k.cdc.Unmarshal(value, &gasBid); err != nil {
			return err
		}

		gasBids = append(gasBids, gasBid)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGasBidResponse{GasBid: gasBids, Pagination: pageRes}, nil
}

func (k Keeper) GasBid(c context.Context, req *types.QueryGetGasBidRequest) (*types.QueryGetGasBidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetGasBid(
		ctx,
		req.Chain,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetGasBidResponse{GasBid: val}, nil
}
