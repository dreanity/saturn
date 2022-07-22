package cli

import (
	"context"

	"saturn/x/randomness/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListUnprovenRandomness() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-unproven-randomness-2",
		Short: "list all UnprovenRandomness",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllUnprovenRandomnessRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.UnprovenRandomnessAll(context.Background(), params)
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

func CmdShowUnprovenRandomness() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-unproven-randomness-2 [round]",
		Short: "shows a UnprovenRandomness",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argRound, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetUnprovenRandomnessRequest{
				Round: argRound,
			}

			res, err := queryClient.UnprovenRandomness(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
