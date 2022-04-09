package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/zarazan/apollo/testutil/keeper"
	"github.com/zarazan/apollo/x/apollo/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ApolloKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
