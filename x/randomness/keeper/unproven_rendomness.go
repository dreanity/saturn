package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"saturn/x/randomness/types"
)

// SetUnprovenRendomness set a specific unprovenRendomness in the store from its index
func (k Keeper) SetUnprovenRendomness(ctx sdk.Context, unprovenRendomness types.UnprovenRendomness) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UnprovenRendomnessKeyPrefix))
	b := k.cdc.MustMarshal(&unprovenRendomness)
	store.Set(types.UnprovenRendomnessKey(
		unprovenRendomness.Index,
	), b)
}

// GetUnprovenRendomness returns a unprovenRendomness from its index
func (k Keeper) GetUnprovenRendomness(
	ctx sdk.Context,
	index string,

) (val types.UnprovenRendomness, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UnprovenRendomnessKeyPrefix))

	b := store.Get(types.UnprovenRendomnessKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUnprovenRendomness removes a unprovenRendomness from the store
func (k Keeper) RemoveUnprovenRendomness(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UnprovenRendomnessKeyPrefix))
	store.Delete(types.UnprovenRendomnessKey(
		index,
	))
}

// GetAllUnprovenRendomness returns all unprovenRendomness
func (k Keeper) GetAllUnprovenRendomness(ctx sdk.Context) (list []types.UnprovenRendomness) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UnprovenRendomnessKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UnprovenRendomness
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
