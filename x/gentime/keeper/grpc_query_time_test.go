package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/dreanity/saturn/testutil/keeper"
	"github.com/dreanity/saturn/testutil/nullify"
	"github.com/dreanity/saturn/x/gentime/types"
)

func TestTimeQuery(t *testing.T) {
	keeper, ctx := keepertest.GentimeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestTime(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTimeRequest
		response *types.QueryGetTimeResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetTimeRequest{},
			response: &types.QueryGetTimeResponse{Time: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Time(wctx, tc.request)
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
