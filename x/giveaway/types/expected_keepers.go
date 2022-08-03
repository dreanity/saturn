package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	genTimeTypes "github.com/dreanity/saturn/x/gentime/types"
	randomnessTypes "github.com/dreanity/saturn/x/randomness/types"
)

type GentimeKeeper interface {
	GetTime(ctx sdk.Context) (val genTimeTypes.Time)
	// Methods imported from bank should be defined here
}

type RandomnessKeeper interface {
	ComputeRandomnessRoundForTime(ctx sdk.Context, time uint64) uint64
	SetUnprovenRandomness(ctx sdk.Context, unprovenRandomness randomnessTypes.UnprovenRandomness)
	GetProvenRandomness(ctx sdk.Context, round uint64) (val randomnessTypes.ProvenRandomness, found bool)
	GetUnprovenRandomness(ctx sdk.Context, round uint64) (val randomnessTypes.UnprovenRandomness, found bool)
	// Methods imported from bank should be defined here
}
