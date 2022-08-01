package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/giveaway/types"
)

// SetGiveaway set a specific giveaway in the store from its index
func (k Keeper) SetGiveaway(ctx sdk.Context, giveaway types.Giveaway) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GiveawayKeyPrefix))
	b := k.cdc.MustMarshal(&giveaway)
	store.Set(types.GiveawayKey(
		giveaway.Index,
	), b)
}

// GetGiveaway returns a giveaway from its index
func (k Keeper) GetGiveaway(
	ctx sdk.Context,
	index uint64,

) (val types.Giveaway, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GiveawayKeyPrefix))

	b := store.Get(types.GiveawayKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGiveaway removes a giveaway from the store
func (k Keeper) RemoveGiveaway(
	ctx sdk.Context,
	index uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GiveawayKeyPrefix))
	store.Delete(types.GiveawayKey(
		index,
	))
}

// GetAllGiveaway returns all giveaway
func (k Keeper) GetAllGiveaway(ctx sdk.Context) (list []types.Giveaway) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GiveawayKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Giveaway
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
