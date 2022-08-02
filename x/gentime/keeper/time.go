package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/gentime/types"
)

// SetTime set time in the store
func (k Keeper) SetTime(ctx sdk.Context, time types.Time) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimeKey))
	b := k.cdc.MustMarshal(&time)
	store.Set([]byte{0}, b)
}

// GetTime returns time
func (k Keeper) GetTime(ctx sdk.Context) (val types.Time) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimeKey))

	b := store.Get([]byte{0})

	k.cdc.MustUnmarshal(b, &val)
	return val
}

// RemoveTime removes time from the store
func (k Keeper) RemoveTime(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimeKey))
	store.Delete([]byte{0})
}
