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

func TestGiveawayByHeightQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNGiveawayByHeight(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetGiveawayByHeightRequest
		response *types.QueryGetGiveawayByHeightResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetGiveawayByHeightRequest{
				Height: msgs[0].Height,
			},
			response: &types.QueryGetGiveawayByHeightResponse{GiveawayByHeight: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetGiveawayByHeightRequest{
				Height: msgs[1].Height,
			},
			response: &types.QueryGetGiveawayByHeightResponse{GiveawayByHeight: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetGiveawayByHeightRequest{
				Height: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GiveawayByHeight(wctx, tc.request)
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

func TestGiveawayByHeightQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.GiveawayKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNGiveawayByHeight(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllGiveawayByHeightRequest {
		return &types.QueryAllGiveawayByHeightRequest{
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
			resp, err := keeper.GiveawayByHeightAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.GiveawayByHeight), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.GiveawayByHeight),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.GiveawayByHeightAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.GiveawayByHeight), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.GiveawayByHeight),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.GiveawayByHeightAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.GiveawayByHeight),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.GiveawayByHeightAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
