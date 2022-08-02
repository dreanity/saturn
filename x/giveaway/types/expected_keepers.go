package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	genTimeTypes "github.com/dreanity/saturn/x/gentime/types"
)

type GentimeKeeper interface {
	GetTime(ctx sdk.Context) (val genTimeTypes.Time)
	// Methods imported from bank should be defined here
}
