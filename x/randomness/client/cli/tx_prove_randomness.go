package cli

import (
	"strconv"

	"saturn/x/randomness/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdProveRandomness() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prove-randomness [round] [randomness] [signature] [previous-signature]",
		Short: "Broadcast message proveRandomness",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRound, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argRandomness := args[1]
			argSignature := args[2]
			argPreviousSignature := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgProveRandomness(
				clientCtx.GetFromAddress().String(),
				argRound,
				argRandomness,
				argSignature,
				argPreviousSignature,
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
