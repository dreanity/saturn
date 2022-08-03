package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/dreanity/saturn/x/giveaway/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListGiveawayByHeight() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-giveaway-by-height",
		Short: "list all giveawayByHeight",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllGiveawayByHeightRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.GiveawayByHeightAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowGiveawayByHeight() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-giveaway-by-height [height]",
		Short: "shows a giveawayByHeight",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argHeight, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetGiveawayByHeightRequest{
				Height: argHeight,
			}

			res, err := queryClient.GiveawayByHeight(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
