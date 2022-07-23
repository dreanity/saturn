package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgProveRandomness = "prove_randomness"

var _ sdk.Msg = &MsgProveRandomness{}

func NewMsgProveRandomness(creator string, round uint64, randomness string, signature string, previousSignature string) *MsgProveRandomness {
	return &MsgProveRandomness{
		Creator:           creator,
		Round:             round,
		Randomness:        randomness,
		Signature:         signature,
		PreviousSignature: previousSignature,
	}
}

func (msg *MsgProveRandomness) Route() string {
	return RouterKey
}

func (msg *MsgProveRandomness) Type() string {
	return TypeMsgProveRandomness
}

func (msg *MsgProveRandomness) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgProveRandomness) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProveRandomness) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
