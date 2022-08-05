package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/giveaway module sentinel errors
var (
	ErrDurationToLow                     = sdkerrors.Register(ModuleName, 1100, "duration must be greater than 3600s")
	ErrNameIsShort                       = sdkerrors.Register(ModuleName, 1101, "length of name must be greater than 0")
	ErrPrizesMustBeMoreThanOne           = sdkerrors.Register(ModuleName, 1102, "prizes length must be more than one")
	ErrPrizeAmountMustBeMoreThanZero     = sdkerrors.Register(ModuleName, 1104, "prize amount must be more than 0")
	ErrPrizeScaleToBig                   = sdkerrors.Register(ModuleName, 1105, "prize scale to big")
	ErrPrizeCurrencyMustBeMoreThanZero   = sdkerrors.Register(ModuleName, 1106, "length of prize currency must be more than zero")
	ErrIssueTicketForNonExistentGiveaway = sdkerrors.Register(ModuleName, 1107, "attempt to issue a ticket for a non-existent giveaway")
	ErrInvalidParticipantId              = sdkerrors.Register(ModuleName, 1108, "participantId must not be empty")
	ErrInvalidParticipantName            = sdkerrors.Register(ModuleName, 1109, "participantName must not be empty")
	ErrTicketRegistrationClosed          = sdkerrors.Register(ModuleName, 1110, "ticket registration closed")
)
