package keeper

import (
	"github.com/dreanity/saturn/x/randomness/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetChainInfo set chainInfo in the store
func (k Keeper) SetChainInfo(ctx sdk.Context, chainInfo types.ChainInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainInfoKey))
	b := k.cdc.MustMarshal(&chainInfo)
	store.Set([]byte{0}, b)

}

// GetChainInfo returns chainInfo
func (k Keeper) GetChainInfo(ctx sdk.Context) types.ChainInfo {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainInfoKey))
	chainInfo := types.ChainInfo{}

	b := store.Get([]byte{0})
	k.cdc.MustUnmarshal(b, &chainInfo)
	return chainInfo
}
