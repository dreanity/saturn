package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/profile module sentinel errors
var (
	ErrSample          = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrNameIsShort     = sdkerrors.Register(ModuleName, 1101, "length of name must be greater than 0")
	ErrNameAlreadyUsed = sdkerrors.Register(ModuleName, 1102, "name already used")
)
