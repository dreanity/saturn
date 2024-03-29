package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/giveaway/types"
)

// SetTicket set a specific ticket in the store from its index
func (k Keeper) SetTicket(ctx sdk.Context, ticket types.Ticket) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TicketKeyPrefix))
	b := k.cdc.MustMarshal(&ticket)
	store.Set(types.TicketKey(
		ticket.GiveawayId,
		ticket.Index,
	), b)
}

// GetTicket returns a ticket from its index
func (k Keeper) GetTicket(
	ctx sdk.Context,
	giveawayId uint32,
	index uint32,

) (val types.Ticket, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TicketKeyPrefix))

	b := store.Get(types.TicketKey(
		giveawayId,
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTicket removes a ticket from the store
func (k Keeper) RemoveTicket(
	ctx sdk.Context,
	giveawayId uint32,
	index uint32,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TicketKeyPrefix))
	store.Delete(types.TicketKey(
		giveawayId,
		index,
	))
}

// GetAllTicket returns all ticket
func (k Keeper) GetAllTicket(ctx sdk.Context) (list []types.Ticket) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TicketKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Ticket
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
