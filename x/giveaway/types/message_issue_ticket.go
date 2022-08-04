package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgIssueTicket = "issue_ticket"

var _ sdk.Msg = &MsgIssueTicket{}

func NewMsgIssueTicket(creator string, giveawayId uint32, participantId string, participantName string) *MsgIssueTicket {
	return &MsgIssueTicket{
		Creator:         creator,
		GiveawayId:      giveawayId,
		ParticipantId:   participantId,
		ParticipantName: participantName,
	}
}

func (msg *MsgIssueTicket) Route() string {
	return RouterKey
}

func (msg *MsgIssueTicket) Type() string {
	return TypeMsgIssueTicket
}

func (msg *MsgIssueTicket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgIssueTicket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgIssueTicket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
