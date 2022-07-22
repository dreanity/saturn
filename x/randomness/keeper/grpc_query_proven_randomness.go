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

func (k Keeper) ProvenRandomnessAll(c context.Context, req *types.QueryAllProvenRandomnessRequest) (*types.QueryAllProvenRandomnessResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var provenRandomnesss []types.ProvenRandomness
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	provenRandomnessStore := prefix.NewStore(store, types.KeyPrefix(types.ProvenRandomnessKeyPrefix))

	pageRes, err := query.Paginate(provenRandomnessStore, req.Pagination, func(key []byte, value []byte) error {
		var provenRandomness types.ProvenRandomness
		if err := k.cdc.Unmarshal(value, &provenRandomness); err != nil {
			return err
		}

		provenRandomnesss = append(provenRandomnesss, provenRandomness)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProvenRandomnessResponse{ProvenRandomness: provenRandomnesss, Pagination: pageRes}, nil
}

func (k Keeper) ProvenRandomness(c context.Context, req *types.QueryGetProvenRandomnessRequest) (*types.QueryGetProvenRandomnessResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetProvenRandomness(
		ctx,
		req.Round,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetProvenRandomnessResponse{ProvenRandomness: val}, nil
}
