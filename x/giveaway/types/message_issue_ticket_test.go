package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dreanity/saturn/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgIssueTicket_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgIssueTicket
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgIssueTicket{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgIssueTicket{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
