package keeper

import (
	"github.com/dreanity/saturn/x/gentime/types"
)

var _ types.QueryServer = Keeper{}
