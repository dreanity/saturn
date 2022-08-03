package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/dreanity/saturn/x/giveaway/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListGiveawayByRandomness() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-giveaway-by-randomness",
		Short: "list all giveawayByRandomness",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllGiveawayByRandomnessRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.GiveawayByRandomnessAll(context.Background(), params)
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

func CmdShowGiveawayByRandomness() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-giveaway-by-randomness [round]",
		Short: "shows a giveawayByRandomness",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argRound, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetGiveawayByRandomnessRequest{
				Round: argRound,
			}

			res, err := queryClient.GiveawayByRandomness(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
