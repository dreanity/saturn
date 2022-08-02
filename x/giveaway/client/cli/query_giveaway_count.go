package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/dreanity/saturn/x/giveaway/types"
	"github.com/spf13/cobra"
)

func CmdShowGiveawayCount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-giveaway-count",
		Short: "shows giveawayCount",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetGiveawayCountRequest{}

			res, err := queryClient.GiveawayCount(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
