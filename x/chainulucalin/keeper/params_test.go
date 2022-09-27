package keeper_test

import (
	"testing"

	testkeeper "chainulucalin/testutil/keeper"
	"chainulucalin/x/chainulucalin/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ChainulucalinKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
