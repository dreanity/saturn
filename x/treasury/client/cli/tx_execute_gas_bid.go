package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/dreanity/saturn/x/treasury/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdExecuteGasBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute-gas-bid [chain] [token-address] [paid-amount] [recipient] [bid-number]",
		Short: "Broadcast message execute_gas_bid",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argChain := args[0]
			argTokenAddress := args[1]
			argPaidAmount := args[2]
			argRecipient := args[3]
			argBidNumber, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgExecuteGasBid(
				clientCtx.GetFromAddress().String(),
				argTokenAddress,
				argPaidAmount,
				argRecipient,
				argBidNumber,
				argChain,
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
