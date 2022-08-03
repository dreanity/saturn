package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/dreanity/saturn/x/giveaway/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GiveawayByRandomnessAll(c context.Context, req *types.QueryAllGiveawayByRandomnessRequest) (*types.QueryAllGiveawayByRandomnessResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var giveawayByRandomnesss []types.GiveawayByRandomness
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	giveawayByRandomnessStore := prefix.NewStore(store, types.KeyPrefix(types.GiveawayByRandomnessKeyPrefix))

	pageRes, err := query.Paginate(giveawayByRandomnessStore, req.Pagination, func(key []byte, value []byte) error {
		var giveawayByRandomness types.GiveawayByRandomness
		if err := k.cdc.Unmarshal(value, &giveawayByRandomness); err != nil {
			return err
		}

		giveawayByRandomnesss = append(giveawayByRandomnesss, giveawayByRandomness)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGiveawayByRandomnessResponse{GiveawayByRandomness: giveawayByRandomnesss, Pagination: pageRes}, nil
}

func (k Keeper) GiveawayByRandomness(c context.Context, req *types.QueryGetGiveawayByRandomnessRequest) (*types.QueryGetGiveawayByRandomnessResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetGiveawayByRandomness(
		ctx,
		req.Round,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetGiveawayByRandomnessResponse{GiveawayByRandomness: val}, nil
}
