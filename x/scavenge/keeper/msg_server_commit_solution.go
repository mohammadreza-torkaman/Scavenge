package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mohammadreza-torkaman/scavenge/x/scavenge/types"
)

func (k msgServer) CommitSolution(goCtx context.Context, msg *types.MsgCommitSolution) (*types.MsgCommitSolutionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetCommit(ctx, msg.SolutionHash)

	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "This commit already exists")
	}
	commit := types.Commit{
		Index:                 msg.SolutionHash,
		SolutionHash:          msg.SolutionHash,
		SolutionScavengerHash: msg.SolutionScavengerHash,
	}
	k.SetCommit(ctx, commit)

	return &types.MsgCommitSolutionResponse{}, nil
}
