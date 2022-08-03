package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/giveaway/types"
)

// SetGiveawayByRandomness set a specific giveawayByRandomness in the store from its index
func (k Keeper) SetGiveawayByRandomness(ctx sdk.Context, giveawayByRandomness types.GiveawayByRandomness) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GiveawayByRandomnessKeyPrefix))
	b := k.cdc.MustMarshal(&giveawayByRandomness)
	store.Set(types.GiveawayByRandomnessKey(
		giveawayByRandomness.Round,
	), b)
}

// GetGiveawayByRandomness returns a giveawayByRandomness from its index
func (k Keeper) GetGiveawayByRandomness(
	ctx sdk.Context,
	round uint64,

) (val types.GiveawayByRandomness, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GiveawayByRandomnessKeyPrefix))

	b := store.Get(types.GiveawayByRandomnessKey(
		round,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGiveawayByRandomness removes a giveawayByRandomness from the store
func (k Keeper) RemoveGiveawayByRandomness(
	ctx sdk.Context,
	round uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GiveawayByRandomnessKeyPrefix))
	store.Delete(types.GiveawayByRandomnessKey(
		round,
	))
}

// GetAllGiveawayByRandomness returns all giveawayByRandomness
func (k Keeper) GetAllGiveawayByRandomness(ctx sdk.Context) (list []types.GiveawayByRandomness) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GiveawayByRandomnessKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.GiveawayByRandomness
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
