package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/dreanity/saturn/x/profile/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NameRegistryAll(c context.Context, req *types.QueryAllNameRegistryRequest) (*types.QueryAllNameRegistryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nameRegistrys []types.NameRegistry
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	nameRegistryStore := prefix.NewStore(store, types.KeyPrefix(types.NameRegistryKeyPrefix))

	pageRes, err := query.Paginate(nameRegistryStore, req.Pagination, func(key []byte, value []byte) error {
		var nameRegistry types.NameRegistry
		if err := k.cdc.Unmarshal(value, &nameRegistry); err != nil {
			return err
		}

		nameRegistrys = append(nameRegistrys, nameRegistry)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNameRegistryResponse{NameRegistry: nameRegistrys, Pagination: pageRes}, nil
}

func (k Keeper) NameRegistry(c context.Context, req *types.QueryGetNameRegistryRequest) (*types.QueryGetNameRegistryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNameRegistry(
		ctx,
		req.Name,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNameRegistryResponse{NameRegistry: val}, nil
}
