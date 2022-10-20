package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/profile/types"
)

// SetNameRegistry set a specific nameRegistry in the store from its index
func (k Keeper) SetNameRegistry(ctx sdk.Context, nameRegistry types.NameRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameRegistryKeyPrefix))
	b := k.cdc.MustMarshal(&nameRegistry)
	store.Set(types.NameRegistryKey(
		nameRegistry.Name,
	), b)
}

// GetNameRegistry returns a nameRegistry from its index
func (k Keeper) GetNameRegistry(
	ctx sdk.Context,
	name string,

) (val types.NameRegistry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameRegistryKeyPrefix))

	b := store.Get(types.NameRegistryKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNameRegistry removes a nameRegistry from the store
func (k Keeper) RemoveNameRegistry(
	ctx sdk.Context,
	name string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameRegistryKeyPrefix))
	store.Delete(types.NameRegistryKey(
		name,
	))
}

// GetAllNameRegistry returns all nameRegistry
func (k Keeper) GetAllNameRegistry(ctx sdk.Context) (list []types.NameRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameRegistryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NameRegistry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
