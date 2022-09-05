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

func (k Keeper) GasPriceAll(c context.Context, req *types.QueryAllGasPriceRequest) (*types.QueryAllGasPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var gasPrices []types.GasPrice
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	gasPriceStore := prefix.NewStore(store, types.KeyPrefix(types.GasPriceKeyPrefix))

	pageRes, err := query.Paginate(gasPriceStore, req.Pagination, func(key []byte, value []byte) error {
		var gasPrice types.GasPrice
		if err := k.cdc.Unmarshal(value, &gasPrice); err != nil {
			return err
		}

		gasPrices = append(gasPrices, gasPrice)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGasPriceResponse{GasPrice: gasPrices, Pagination: pageRes}, nil
}

func (k Keeper) GasPrice(c context.Context, req *types.QueryGetGasPriceRequest) (*types.QueryGetGasPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetGasPrice(
		ctx,
		req.Chain,
		req.TokenAddress,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetGasPriceResponse{GasPrice: val}, nil
}
