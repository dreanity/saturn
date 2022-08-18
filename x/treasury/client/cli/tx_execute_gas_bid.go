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
		Use:   "execute-gas-bid [currency] [paid-amount] [recipient] [bid-number] [from-chain]",
		Short: "Broadcast message execute_gas_bid",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCurrency := args[0]
			argPaidAmount := args[1]
			argRecipient := args[2]
			argBidNumber, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			argFromChain := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgExecuteGasBid(
				clientCtx.GetFromAddress().String(),
				argCurrency,
				argPaidAmount,
				argRecipient,
				argBidNumber,
				argFromChain,
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
