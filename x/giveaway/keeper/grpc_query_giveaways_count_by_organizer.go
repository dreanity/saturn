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

func (k Keeper) GiveawaysCountByOrganizerAll(c context.Context, req *types.QueryAllGiveawaysCountByOrganizerRequest) (*types.QueryAllGiveawaysCountByOrganizerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var giveawaysCountByOrganizers []types.GiveawaysCountByOrganizer
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	giveawaysCountByOrganizerStore := prefix.NewStore(store, types.KeyPrefix(types.GiveawaysCountByOrganizerKeyPrefix))

	pageRes, err := query.Paginate(giveawaysCountByOrganizerStore, req.Pagination, func(key []byte, value []byte) error {
		var giveawaysCountByOrganizer types.GiveawaysCountByOrganizer
		if err := k.cdc.Unmarshal(value, &giveawaysCountByOrganizer); err != nil {
			return err
		}

		giveawaysCountByOrganizers = append(giveawaysCountByOrganizers, giveawaysCountByOrganizer)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGiveawaysCountByOrganizerResponse{GiveawaysCountByOrganizer: giveawaysCountByOrganizers, Pagination: pageRes}, nil
}

func (k Keeper) GiveawaysCountByOrganizer(c context.Context, req *types.QueryGetGiveawaysCountByOrganizerRequest) (*types.QueryGetGiveawaysCountByOrganizerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetGiveawaysCountByOrganizer(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetGiveawaysCountByOrganizerResponse{GiveawaysCountByOrganizer: val}, nil
}
