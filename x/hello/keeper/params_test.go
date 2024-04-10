package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "hello/testutil/keeper"
	"hello/x/hello/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.HelloKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
