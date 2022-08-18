package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/dreanity/saturn/x/treasury/types"
	"github.com/spf13/cobra"
)

func CmdListGasPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-gas-price",
		Short: "list all gas_price",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllGasPriceRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.GasPriceAll(context.Background(), params)
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

func CmdShowGasPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-gas-price [currency]",
		Short: "shows a gas_price",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argCurrency := args[0]

			params := &types.QueryGetGasPriceRequest{
				Currency: argCurrency,
			}

			res, err := queryClient.GasPrice(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
