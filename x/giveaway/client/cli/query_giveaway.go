package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/dreanity/saturn/x/giveaway/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListGiveaway() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-giveaway",
		Short: "list all giveaway",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllGiveawayRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.GiveawayAll(context.Background(), params)
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

func CmdShowGiveaway() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-giveaway [index]",
		Short: "shows a giveaway",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetGiveawayRequest{
				Index: argIndex,
			}

			res, err := queryClient.Giveaway(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
