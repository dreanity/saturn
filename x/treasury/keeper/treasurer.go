package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/treasury/types"
)

// SetTreasurer set treasurer in the store
func (k Keeper) SetTreasurer(ctx sdk.Context, treasurer types.Treasurer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TreasurerKey))
	b := k.cdc.MustMarshal(&treasurer)
	store.Set([]byte{0}, b)
}

// GetTreasurer returns treasurer
func (k Keeper) GetTreasurer(ctx sdk.Context) (val types.Treasurer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TreasurerKey))

	b := store.Get([]byte{0})

	k.cdc.MustUnmarshal(b, &val)
	return val
}
