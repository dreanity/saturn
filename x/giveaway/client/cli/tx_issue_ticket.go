package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/dreanity/saturn/x/giveaway/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdIssueTicket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue-ticket [giveaway-id] [participant-id] [participant-name]",
		Short: "Issue a ticket to distribution participant",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argGiveawayId, err := cast.ToUint32E(args[0])
			if err != nil {
				return err
			}
			argParticipantId := args[1]
			argParticipantName := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgIssueTicket(
				clientCtx.GetFromAddress().String(),
				argGiveawayId,
				argParticipantId,
				argParticipantName,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
