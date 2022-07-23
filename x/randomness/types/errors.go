package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/randomness module sentinel errors
var (
	ErrUnprovenRandomnessNotExists = sdkerrors.Register(ModuleName, 1100, "unproven randomnness not exists")
	ErrRandomnessVerification      = sdkerrors.Register(ModuleName, 1101, "randomnness verification failed")
)
