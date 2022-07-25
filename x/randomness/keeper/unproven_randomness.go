package keeper

import (
	"github.com/dreanity/saturn/x/randomness/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ComputeRandomnessRoundForTime(ctx sdk.Context, time uint64) uint64 {
	chainInfo := k.GetChainInfo(ctx)
	round := uint64(1)

	if time < chainInfo.GenesisTime {
		round = 1
	} else {
		timeDiff := sdk.NewDecFromBigInt(sdk.NewIntFromUint64(time - chainInfo.GenesisTime).BigInt())
		roundForTime := timeDiff.QuoRoundUp(sdk.NewDecFromInt(sdk.NewIntFromUint64(chainInfo.Period)))
		nextRound := roundForTime.Ceil().TruncateInt().Uint64() + 1

		round = nextRound
	}

	return round
}

// SetUnprovenRandomness set a specific unprovenRandomness in the store from its index
func (k Keeper) SetUnprovenRandomness(ctx sdk.Context, unprovenRandomness types.UnprovenRandomness) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UnprovenRandomnessKeyPrefix))
	b := k.cdc.MustMarshal(&unprovenRandomness)
	store.Set(types.UnprovenRandomnessKey(
		unprovenRandomness.Round,
	), b)
}

// GetUnprovenRandomness returns a unprovenRandomness from its index
func (k Keeper) GetUnprovenRandomness(
	ctx sdk.Context,
	round uint64,

) (val types.UnprovenRandomness, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UnprovenRandomnessKeyPrefix))

	b := store.Get(types.UnprovenRandomnessKey(
		round,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUnprovenRandomness removes a unprovenRandomness from the store
func (k Keeper) RemoveUnprovenRandomness(
	ctx sdk.Context,
	round uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UnprovenRandomnessKeyPrefix))
	store.Delete(types.UnprovenRandomnessKey(
		round,
	))
}

// GetAllUnprovenRandomness returns all unprovenRandomness
func (k Keeper) GetAllUnprovenRandomness(ctx sdk.Context) (list []types.UnprovenRandomness) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UnprovenRandomnessKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UnprovenRandomness
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
