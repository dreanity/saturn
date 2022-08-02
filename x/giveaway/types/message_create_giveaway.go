package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateGiveaway = "create_giveaway"

var _ sdk.Msg = &MsgCreateGiveaway{}

func NewMsgCreateGiveaway(creator string, duration int32, roundCount uint64, name string, prize *Prize) *MsgCreateGiveaway {
	return &MsgCreateGiveaway{
		Creator:    creator,
		Duration:   duration,
		RoundCount: roundCount,
		Name:       name,
		Prize:      prize,
	}
}

func (msg *MsgCreateGiveaway) Route() string {
	return RouterKey
}

func (msg *MsgCreateGiveaway) Type() string {
	return TypeMsgCreateGiveaway
}

func (msg *MsgCreateGiveaway) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateGiveaway) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateGiveaway) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
