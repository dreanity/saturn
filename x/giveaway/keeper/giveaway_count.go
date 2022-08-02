package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/giveaway/types"
)

// SetGiveawayCount set giveawayCount in the store
func (k Keeper) SetGiveawayCount(ctx sdk.Context, giveawayCount types.GiveawayCount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GiveawayCountKey))
	b := k.cdc.MustMarshal(&giveawayCount)
	store.Set([]byte{0}, b)
}

// GetGiveawayCount returns giveawayCount
func (k Keeper) GetGiveawayCount(ctx sdk.Context) (val types.GiveawayCount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GiveawayCountKey))

	b := store.Get([]byte{0})

	k.cdc.MustUnmarshal(b, &val)
	return val
}

// RemoveGiveawayCount removes giveawayCount from the store
func (k Keeper) RemoveGiveawayCount(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GiveawayCountKey))
	store.Delete([]byte{0})
}
