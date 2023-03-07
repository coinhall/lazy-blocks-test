package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/coinhall/test/testutil/keeper"
	"github.com/coinhall/test/x/test/keeper"
	"github.com/coinhall/test/x/test/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TestKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
