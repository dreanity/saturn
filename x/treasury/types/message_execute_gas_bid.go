package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgExecuteGasBid = "execute_gas_bid"

var _ sdk.Msg = &MsgExecuteGasBid{}

func NewMsgExecuteGasBid(creator string, tokenAddress string, paidAmount string, recipient string, bidNumber uint64, chain string) *MsgExecuteGasBid {
	return &MsgExecuteGasBid{
		Creator:      creator,
		TokenAddress: tokenAddress,
		PaidAmount:   paidAmount,
		Recipient:    recipient,
		BidNumber:    bidNumber,
		Chain:        chain,
	}
}

func (msg *MsgExecuteGasBid) Route() string {
	return RouterKey
}

func (msg *MsgExecuteGasBid) Type() string {
	return TypeMsgExecuteGasBid
}

func (msg *MsgExecuteGasBid) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgExecuteGasBid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgExecuteGasBid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
