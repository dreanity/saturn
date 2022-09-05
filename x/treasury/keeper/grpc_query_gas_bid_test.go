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
	"github.com/dreanity/saturn/x/treasury/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestGasBidQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.TreasuryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNGasBid(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetGasBidRequest
		response *types.QueryGetGasBidResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetGasBidRequest{
				Chain: msgs[0].Chain,
			},
			response: &types.QueryGetGasBidResponse{GasBid: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetGasBidRequest{
				Chain: msgs[1].Chain,
			},
			response: &types.QueryGetGasBidResponse{GasBid: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetGasBidRequest{
				Chain: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GasBid(wctx, tc.request)
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

func TestGasBidQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.TreasuryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNGasBid(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllGasBidRequest {
		return &types.QueryAllGasBidRequest{
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
			resp, err := keeper.GasBidAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.GasBid), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.GasBid),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.GasBidAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.GasBid), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.GasBid),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.GasBidAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.GasBid),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.GasBidAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
