package chainulucalin_test

import (
	"testing"

	keepertest "chainulucalin/testutil/keeper"
	"chainulucalin/testutil/nullify"
	"chainulucalin/x/chainulucalin"
	"chainulucalin/x/chainulucalin/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ChainulucalinKeeper(t)
	chainulucalin.InitGenesis(ctx, *k, genesisState)
	got := chainulucalin.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
