package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/orionis-labs/orionis/x/orionis/types"
)

func (k msgServer) OnrInr(goCtx context.Context, msg *types.MsgOnrInr) (*types.MsgOnrInrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgOnrInrResponse{}, nil
}
