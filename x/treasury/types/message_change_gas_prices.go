package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgChangeGasPrices = "change_gas_prices"

var _ sdk.Msg = &MsgChangeGasPrices{}

func NewMsgChangeGasPrices(creator string, gasPrices []*GasPrice) *MsgChangeGasPrices {
	return &MsgChangeGasPrices{
		Creator:   creator,
		GasPrices: gasPrices,
	}
}

func (msg *MsgChangeGasPrices) Route() string {
	return RouterKey
}

func (msg *MsgChangeGasPrices) Type() string {
	return TypeMsgChangeGasPrices
}

func (msg *MsgChangeGasPrices) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgChangeGasPrices) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgChangeGasPrices) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
