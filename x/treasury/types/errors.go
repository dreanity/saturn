package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/treasury module sentinel errors
var (
	ErrInvalidGasPrice       = sdkerrors.Register(ModuleName, 1101, "invalid gas price")
	ErrCreatorIsNotTreasurer = sdkerrors.Register(ModuleName, 1102, "the creator is not the treasurer")
)
