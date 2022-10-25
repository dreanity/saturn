package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/dreanity/saturn/x/giveaway/types"
	"github.com/spf13/cobra"
)

func CmdListGiveawaysCountByOrganizer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-giveaways-count-by-organizer",
		Short: "list all giveaways_count_by_organizer",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllGiveawaysCountByOrganizerRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.GiveawaysCountByOrganizerAll(context.Background(), params)
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

func CmdShowGiveawaysCountByOrganizer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-giveaways-count-by-organizer [address]",
		Short: "shows a giveaways_count_by_organizer",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]

			params := &types.QueryGetGiveawaysCountByOrganizerRequest{
				Address: argAddress,
			}

			res, err := queryClient.GiveawaysCountByOrganizer(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
