package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/treasury/types"
)

// SetGasBid set a specific gasBid in the store from its index
func (k Keeper) SetGasBid(ctx sdk.Context, gasBid types.GasBid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GasBidKeyPrefix))
	b := k.cdc.MustMarshal(&gasBid)
	store.Set(types.GasBidKey(
		gasBid.FromChain,
	), b)
}

// GetGasBid returns a gasBid from its index
func (k Keeper) GetGasBid(
	ctx sdk.Context,
	fromChain string,

) (val types.GasBid, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GasBidKeyPrefix))

	b := store.Get(types.GasBidKey(
		fromChain,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGasBid removes a gasBid from the store
func (k Keeper) RemoveGasBid(
	ctx sdk.Context,
	fromChain string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GasBidKeyPrefix))
	store.Delete(types.GasBidKey(
		fromChain,
	))
}

// GetAllGasBid returns all gasBid
func (k Keeper) GetAllGasBid(ctx sdk.Context) (list []types.GasBid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GasBidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.GasBid
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
