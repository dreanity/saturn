package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"saturn/x/randomness/types"
)

var _ = strconv.Itoa(0)

func CmdProveRandomness() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prove-randomness [round] [time] [randomness] [signature] [previous-signature]",
		Short: "Broadcast message proveRandomness",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRound, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argTime, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argRandomness := args[2]
			argSignature := args[3]
			argPreviousSignature := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgProveRandomness(
				clientCtx.GetFromAddress().String(),
				argRound,
				argTime,
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
