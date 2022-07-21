package keeper

import (
	"saturn/x/randomness/types"
)

var _ types.QueryServer = Keeper{}
