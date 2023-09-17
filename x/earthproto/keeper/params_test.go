package keeper_test

import (
	"testing"

	testkeeper "earthProto/testutil/keeper"
	"earthProto/x/earthproto/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EarthprotoKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
