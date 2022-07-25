package cli

import (
	"context"

	"github.com/dreanity/saturn/x/randomness/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListProvenRandomness() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-proven-randomness-2",
		Short: "list all ProvenRandomness",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllProvenRandomnessRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ProvenRandomnessAll(context.Background(), params)
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

func CmdShowProvenRandomness() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-proven-randomness-2 [round]",
		Short: "shows a ProvenRandomness",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argRound, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetProvenRandomnessRequest{
				Round: argRound,
			}

			res, err := queryClient.ProvenRandomness(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
