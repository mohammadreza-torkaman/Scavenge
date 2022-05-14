package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mohammadreza-torkaman/scavenge/testutil/keeper"
	"github.com/mohammadreza-torkaman/scavenge/x/scavenge/keeper"
	"github.com/mohammadreza-torkaman/scavenge/x/scavenge/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ScavengeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
