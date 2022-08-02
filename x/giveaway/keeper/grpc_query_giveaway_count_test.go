package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/dreanity/saturn/testutil/keeper"
	"github.com/dreanity/saturn/testutil/nullify"
	"github.com/dreanity/saturn/x/giveaway/types"
)

func TestGiveawayCountQuery(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestGiveawayCount(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetGiveawayCountRequest
		response *types.QueryGetGiveawayCountResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetGiveawayCountRequest{},
			response: &types.QueryGetGiveawayCountResponse{GiveawayCount: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GiveawayCount(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
