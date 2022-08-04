package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/giveaway/types"
)

// SetTicketCount set a specific ticketCount in the store from its index
func (k Keeper) SetTicketCount(ctx sdk.Context, ticketCount types.TicketCount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TicketCountKeyPrefix))
	b := k.cdc.MustMarshal(&ticketCount)
	store.Set(types.TicketCountKey(
		ticketCount.GiveawayId,
	), b)
}

// GetTicketCount returns a ticketCount from its index
func (k Keeper) GetTicketCount(
	ctx sdk.Context,
	giveawayId uint32,

) (val types.TicketCount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TicketCountKeyPrefix))

	b := store.Get(types.TicketCountKey(
		giveawayId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTicketCount removes a ticketCount from the store
func (k Keeper) RemoveTicketCount(
	ctx sdk.Context,
	giveawayId uint32,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TicketCountKeyPrefix))
	store.Delete(types.TicketCountKey(
		giveawayId,
	))
}

// GetAllTicketCount returns all ticketCount
func (k Keeper) GetAllTicketCount(ctx sdk.Context) (list []types.TicketCount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TicketCountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TicketCount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
