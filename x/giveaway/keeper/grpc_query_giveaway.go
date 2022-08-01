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

func (k Keeper) GiveawayAll(c context.Context, req *types.QueryAllGiveawayRequest) (*types.QueryAllGiveawayResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var giveaways []types.Giveaway
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	giveawayStore := prefix.NewStore(store, types.KeyPrefix(types.GiveawayKeyPrefix))

	pageRes, err := query.Paginate(giveawayStore, req.Pagination, func(key []byte, value []byte) error {
		var giveaway types.Giveaway
		if err := k.cdc.Unmarshal(value, &giveaway); err != nil {
			return err
		}

		giveaways = append(giveaways, giveaway)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGiveawayResponse{Giveaway: giveaways, Pagination: pageRes}, nil
}

func (k Keeper) Giveaway(c context.Context, req *types.QueryGetGiveawayRequest) (*types.QueryGetGiveawayResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetGiveaway(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetGiveawayResponse{Giveaway: val}, nil
}
