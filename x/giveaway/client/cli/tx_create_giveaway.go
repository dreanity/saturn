package cli

import (
	"strconv"

	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/dreanity/saturn/x/giveaway/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateGiveaway() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-giveaway [duration] [name] [prizes]",
		Short: "Broadcast message createGiveaway",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDuration, err := cast.ToInt64E(args[0])
			if err != nil {
				return err
			}

			argName := args[1]
			argPrizes := new([]types.Prize)
			err = json.Unmarshal([]byte(args[2]), argPrizes)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateGiveaway(
				clientCtx.GetFromAddress().String(),
				argDuration,
				argName,
				*argPrizes,
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
