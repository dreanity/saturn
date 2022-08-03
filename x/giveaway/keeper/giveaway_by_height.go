package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/giveaway/types"
)

// SetGiveawayByHeight set a specific giveawayByHeight in the store from its index
func (k Keeper) SetGiveawayByHeight(ctx sdk.Context, giveawayByHeight types.GiveawayByHeight) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GiveawayByHeightKeyPrefix))
	b := k.cdc.MustMarshal(&giveawayByHeight)
	store.Set(types.GiveawayByHeightKey(
		giveawayByHeight.Height,
	), b)
}

// GetGiveawayByHeight returns a giveawayByHeight from its index
func (k Keeper) GetGiveawayByHeight(
	ctx sdk.Context,
	height int32,

) (val types.GiveawayByHeight, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GiveawayByHeightKeyPrefix))

	b := store.Get(types.GiveawayByHeightKey(
		height,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGiveawayByHeight removes a giveawayByHeight from the store
func (k Keeper) RemoveGiveawayByHeight(
	ctx sdk.Context,
	height int32,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GiveawayByHeightKeyPrefix))
	store.Delete(types.GiveawayByHeightKey(
		height,
	))
}

// GetAllGiveawayByHeight returns all giveawayByHeight
func (k Keeper) GetAllGiveawayByHeight(ctx sdk.Context) (list []types.GiveawayByHeight) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GiveawayByHeightKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.GiveawayByHeight
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
