package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/dreanity/saturn/x/giveaway/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListTicket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-ticket",
		Short: "list all ticket",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTicketRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.TicketAll(context.Background(), params)
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

func CmdShowTicket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-ticket [giveawayId] [index]",
		Short: "shows a ticket",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex, err := cast.ToUint32E(args[1])
			if err != nil {
				return err
			}

			argGiveawayId, err := cast.ToUint32E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetTicketRequest{
				GiveawayId: argGiveawayId,
				Index:      argIndex,
			}

			res, err := queryClient.Ticket(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
