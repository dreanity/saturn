package keeper

import (
	"github.com/dreanity/saturn/x/randomness/types"
)

var _ types.QueryServer = Keeper{}
