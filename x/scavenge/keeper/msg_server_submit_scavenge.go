package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mohammadreza-torkaman/scavenge/x/scavenge/types"
)

func (k msgServer) SubmitScavenge(goCtx context.Context, msg *types.MsgSubmitScavenge) (*types.MsgSubmitScavengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	scavenger := types.Scavenge{
		Index:        msg.SolutionHash,
		SolutionHash: msg.SolutionHash,
		Description:  msg.Description,
		Reward:       msg.Reward,
		Scavenger:    msg.Creator,
	}
	_, found := k.GetScavenge(ctx, msg.SolutionHash)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Scavenge with this answer already exists")
	}

	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	scavengerAccountAddress, err := sdk.AccAddressFromBech32(scavenger.Scavenger)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Scavenger Account Address invalid")
	}

	reward, parsErr := sdk.ParseCoinsNormalized(scavenger.Reward)

	if parsErr != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Coin parse error")
	}
	sendErr := k.bankKeeper.SendCoins(ctx, scavengerAccountAddress, moduleAcct, reward)
	if sendErr != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Send coins error")
	}
	k.SetScavenge(ctx, scavenger)

	return &types.MsgSubmitScavengeResponse{}, nil
}
