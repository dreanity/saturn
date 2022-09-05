package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/treasury module sentinel errors
var (
	ErrInvalidGasPrice       = sdkerrors.Register(ModuleName, 1101, "invalid gas price")
	ErrCreatorIsNotTreasurer = sdkerrors.Register(ModuleName, 1102, "the creator is not the treasurer")
	ErrInvalidChain          = sdkerrors.Register(ModuleName, 1103, "invalid chain")
	ErrInvalidTokenAddress   = sdkerrors.Register(ModuleName, 1104, "invalid token address")
	ErrInvalidPaidAmount     = sdkerrors.Register(ModuleName, 1105, "invalid paid amount")
	ErrInvalidRecipient      = sdkerrors.Register(ModuleName, 1106, "invalid recipient")
	ErrInvalidBidNumber      = sdkerrors.Register(ModuleName, 1107, "invalid bid number")
	ErrNotFoundGasPrice      = sdkerrors.Register(ModuleName, 1108, "not found gas price")
)
