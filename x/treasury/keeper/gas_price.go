package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/treasury/types"
)

// SetGasPrice set a specific gasPrice in the store from its index
func (k Keeper) SetGasPrice(ctx sdk.Context, gasPrice types.GasPrice) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GasPriceKeyPrefix))
	b := k.cdc.MustMarshal(&gasPrice)
	store.Set(types.GasPriceKey(
		gasPrice.Currency,
	), b)
}

// GetGasPrice returns a gasPrice from its index
func (k Keeper) GetGasPrice(
	ctx sdk.Context,
	currency string,

) (val types.GasPrice, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GasPriceKeyPrefix))

	b := store.Get(types.GasPriceKey(
		currency,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGasPrice removes a gasPrice from the store
func (k Keeper) RemoveGasPrice(
	ctx sdk.Context,
	currency string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GasPriceKeyPrefix))
	store.Delete(types.GasPriceKey(
		currency,
	))
}

// GetAllGasPrice returns all gasPrice
func (k Keeper) GetAllGasPrice(ctx sdk.Context) (list []types.GasPrice) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GasPriceKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.GasPrice
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
