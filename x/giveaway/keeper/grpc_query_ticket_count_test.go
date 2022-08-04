package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/dreanity/saturn/testutil/keeper"
	"github.com/dreanity/saturn/testutil/nullify"
	"github.com/dreanity/saturn/x/giveaway/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestTicketCountQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNTicketCount(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTicketCountRequest
		response *types.QueryGetTicketCountResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetTicketCountRequest{
				GiveawayId: msgs[0].GiveawayId,
			},
			response: &types.QueryGetTicketCountResponse{TicketCount: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetTicketCountRequest{
				GiveawayId: msgs[1].GiveawayId,
			},
			response: &types.QueryGetTicketCountResponse{TicketCount: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetTicketCountRequest{
				GiveawayId: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.TicketCount(wctx, tc.request)
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

func TestTicketCountQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNTicketCount(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllTicketCountRequest {
		return &types.QueryAllTicketCountRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.TicketCountAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.TicketCount), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.TicketCount),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.TicketCountAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.TicketCount), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.TicketCount),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.TicketCountAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.TicketCount),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.TicketCountAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
