package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/giveaway/types"
)

// SetGiveawaysCountByOrganizer set a specific giveawaysCountByOrganizer in the store from its index
func (k Keeper) SetGiveawaysCountByOrganizer(ctx sdk.Context, giveawaysCountByOrganizer types.GiveawaysCountByOrganizer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GiveawaysCountByOrganizerKeyPrefix))
	b := k.cdc.MustMarshal(&giveawaysCountByOrganizer)
	store.Set(types.GiveawaysCountByOrganizerKey(
		giveawaysCountByOrganizer.Address,
	), b)
}

// GetGiveawaysCountByOrganizer returns a giveawaysCountByOrganizer from its index
func (k Keeper) GetGiveawaysCountByOrganizer(
	ctx sdk.Context,
	address string,

) (val types.GiveawaysCountByOrganizer, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GiveawaysCountByOrganizerKeyPrefix))

	b := store.Get(types.GiveawaysCountByOrganizerKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGiveawaysCountByOrganizer removes a giveawaysCountByOrganizer from the store
func (k Keeper) RemoveGiveawaysCountByOrganizer(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GiveawaysCountByOrganizerKeyPrefix))
	store.Delete(types.GiveawaysCountByOrganizerKey(
		address,
	))
}

// GetAllGiveawaysCountByOrganizer returns all giveawaysCountByOrganizer
func (k Keeper) GetAllGiveawaysCountByOrganizer(ctx sdk.Context) (list []types.GiveawaysCountByOrganizer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GiveawaysCountByOrganizerKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.GiveawaysCountByOrganizer
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
