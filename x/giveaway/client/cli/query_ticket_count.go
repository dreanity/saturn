package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/dreanity/saturn/x/giveaway/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListTicketCount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-ticket-count",
		Short: "list all ticket_count",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTicketCountRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.TicketCountAll(context.Background(), params)
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

func CmdShowTicketCount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-ticket-count [giveaway-id]",
		Short: "shows a ticket_count",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argGiveawayId, err := cast.ToUint32E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetTicketCountRequest{
				GiveawayId: argGiveawayId,
			}

			res, err := queryClient.TicketCount(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
