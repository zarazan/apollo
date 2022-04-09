package apollo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zarazan/apollo/testutil/keeper"
	"github.com/zarazan/apollo/testutil/nullify"
	"github.com/zarazan/apollo/x/apollo"
	"github.com/zarazan/apollo/x/apollo/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ApolloKeeper(t)
	apollo.InitGenesis(ctx, *k, genesisState)
	got := apollo.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
