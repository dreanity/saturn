package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/dreanity/saturn/x/treasury/types"
	"github.com/spf13/cobra"
)

func CmdListGasBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-gas-bid",
		Short: "list all gas_bid",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllGasBidRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.GasBidAll(context.Background(), params)
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

func CmdShowGasBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-gas-bid [from-chain]",
		Short: "shows a gas_bid",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argFromChain := args[0]

			params := &types.QueryGetGasBidRequest{
				FromChain: argFromChain,
			}

			res, err := queryClient.GasBid(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
