package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "saturn/testutil/keeper"
	"saturn/testutil/nullify"
	"saturn/x/randomness/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestProvenRandomnessQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RandomnessKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNProvenRandomness(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetProvenRandomnessRequest
		response *types.QueryGetProvenRandomnessResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetProvenRandomnessRequest{
				Round: msgs[0].Round,
			},
			response: &types.QueryGetProvenRandomnessResponse{ProvenRandomness: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetProvenRandomnessRequest{
				Round: msgs[1].Round,
			},
			response: &types.QueryGetProvenRandomnessResponse{ProvenRandomness: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetProvenRandomnessRequest{
				Round: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ProvenRandomness(wctx, tc.request)
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

func TestProvenRandomnessQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RandomnessKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNProvenRandomness(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllProvenRandomnessRequest {
		return &types.QueryAllProvenRandomnessRequest{
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
			resp, err := keeper.ProvenRandomnessAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ProvenRandomness), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ProvenRandomness),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ProvenRandomnessAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ProvenRandomness), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ProvenRandomness),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ProvenRandomnessAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ProvenRandomness),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ProvenRandomnessAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
