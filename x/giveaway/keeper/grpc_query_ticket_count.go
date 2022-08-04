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

func (k Keeper) TicketCountAll(c context.Context, req *types.QueryAllTicketCountRequest) (*types.QueryAllTicketCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ticketCounts []types.TicketCount
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	ticketCountStore := prefix.NewStore(store, types.KeyPrefix(types.TicketCountKeyPrefix))

	pageRes, err := query.Paginate(ticketCountStore, req.Pagination, func(key []byte, value []byte) error {
		var ticketCount types.TicketCount
		if err := k.cdc.Unmarshal(value, &ticketCount); err != nil {
			return err
		}

		ticketCounts = append(ticketCounts, ticketCount)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTicketCountResponse{TicketCount: ticketCounts, Pagination: pageRes}, nil
}

func (k Keeper) TicketCount(c context.Context, req *types.QueryGetTicketCountRequest) (*types.QueryGetTicketCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTicketCount(
		ctx,
		req.GiveawayId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTicketCountResponse{TicketCount: val}, nil
}
