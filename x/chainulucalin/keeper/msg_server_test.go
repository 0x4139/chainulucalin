package keeper_test

import (
	"context"
	"testing"

	keepertest "chainulucalin/testutil/keeper"
	"chainulucalin/x/chainulucalin/keeper"
	"chainulucalin/x/chainulucalin/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ChainulucalinKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
