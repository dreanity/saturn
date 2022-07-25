package keeper

import (
	"github.com/dreanity/saturn/x/randomness/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetProvenRandomness set a specific provenRandomness in the store from its index
func (k Keeper) SetProvenRandomness(ctx sdk.Context, provenRandomness types.ProvenRandomness) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProvenRandomnessKeyPrefix))
	b := k.cdc.MustMarshal(&provenRandomness)
	store.Set(types.ProvenRandomnessKey(
		provenRandomness.Round,
	), b)
}

// GetProvenRandomness returns a provenRandomness from its index
func (k Keeper) GetProvenRandomness(
	ctx sdk.Context,
	round uint64,

) (val types.ProvenRandomness, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProvenRandomnessKeyPrefix))

	b := store.Get(types.ProvenRandomnessKey(
		round,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProvenRandomness removes a provenRandomness from the store
func (k Keeper) RemoveProvenRandomness(
	ctx sdk.Context,
	round uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProvenRandomnessKeyPrefix))
	store.Delete(types.ProvenRandomnessKey(
		round,
	))
}

// GetAllProvenRandomness returns all provenRandomness
func (k Keeper) GetAllProvenRandomness(ctx sdk.Context) (list []types.ProvenRandomness) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProvenRandomnessKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ProvenRandomness
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
