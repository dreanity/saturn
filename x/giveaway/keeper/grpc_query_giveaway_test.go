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

func TestGiveawayQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNGiveaway(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetGiveawayRequest
		response *types.QueryGetGiveawayResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetGiveawayRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetGiveawayResponse{Giveaway: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetGiveawayRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetGiveawayResponse{Giveaway: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetGiveawayRequest{
				Index: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Giveaway(wctx, tc.request)
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

func TestGiveawayQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNGiveaway(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllGiveawayRequest {
		return &types.QueryAllGiveawayRequest{
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
			resp, err := keeper.GiveawayAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Giveaway), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Giveaway),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.GiveawayAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Giveaway), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Giveaway),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.GiveawayAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Giveaway),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.GiveawayAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
