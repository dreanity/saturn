package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"saturn/x/randomness/types"
)

func CmdListUnprovenRendomness() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-unproven-rendomness",
		Short: "list all unproven_rendomness",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllUnprovenRendomnessRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.UnprovenRendomnessAll(context.Background(), params)
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

func CmdShowUnprovenRendomness() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-unproven-rendomness [index]",
		Short: "shows a unproven_rendomness",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetUnprovenRendomnessRequest{
				Index: argIndex,
			}

			res, err := queryClient.UnprovenRendomness(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
