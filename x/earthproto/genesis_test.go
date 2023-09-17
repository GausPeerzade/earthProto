package earthproto_test

import (
	"testing"

	keepertest "earthProto/testutil/keeper"
	"earthProto/testutil/nullify"
	"earthProto/x/earthproto"
	"earthProto/x/earthproto/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EarthprotoKeeper(t)
	earthproto.InitGenesis(ctx, *k, genesisState)
	got := earthproto.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
