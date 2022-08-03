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

func (k Keeper) GiveawayByHeightAll(c context.Context, req *types.QueryAllGiveawayByHeightRequest) (*types.QueryAllGiveawayByHeightResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var giveawayByHeights []types.GiveawayByHeight
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	giveawayByHeightStore := prefix.NewStore(store, types.KeyPrefix(types.GiveawayByHeightKeyPrefix))

	pageRes, err := query.Paginate(giveawayByHeightStore, req.Pagination, func(key []byte, value []byte) error {
		var giveawayByHeight types.GiveawayByHeight
		if err := k.cdc.Unmarshal(value, &giveawayByHeight); err != nil {
			return err
		}

		giveawayByHeights = append(giveawayByHeights, giveawayByHeight)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGiveawayByHeightResponse{GiveawayByHeight: giveawayByHeights, Pagination: pageRes}, nil
}

func (k Keeper) GiveawayByHeight(c context.Context, req *types.QueryGetGiveawayByHeightRequest) (*types.QueryGetGiveawayByHeightResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetGiveawayByHeight(
		ctx,
		req.Height,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetGiveawayByHeightResponse{GiveawayByHeight: val}, nil
}
