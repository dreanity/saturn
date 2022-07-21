package cli

import (
	"context"

	"saturn/x/randomness/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListUnprovenRandomness() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-unproven-randomness",
		Short: "list all unproven_randomness",
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
		Use:   "show-unproven-randomness [index]",
		Short: "shows a unproven_randomness",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetUnprovenRandomnessRequest{
				Index: argIndex,
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
