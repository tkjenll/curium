package keeper

import (
	"context"
	"fmt"

	"github.com/bluzelle/curium/x/crud/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateCrudValue(goCtx context.Context, msg *types.MsgCreateCrudValue) (*types.MsgCreateCrudValueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	if k.HasCrudValue(ctx, msg.Uuid, msg.Key) {
		return nil, sdkerrors.New("crud", 0, "key already exists")
	}

	k.AppendCrudValue(
		ctx,
		msg.Creator,
		msg.Uuid,
		msg.Key,
		msg.Value,
		msg.Lease,
		ctx.BlockHeight(),
	)

	return &types.MsgCreateCrudValueResponse{}, nil
}

func (k msgServer) UpdateCrudValue(goCtx context.Context, msg *types.MsgUpdateCrudValue) (*types.MsgUpdateCrudValueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var CrudValue = types.CrudValue{
		Creator: msg.Creator,
		Uuid:    msg.Uuid,
		Key:     msg.Key,
		Value:   msg.Value,
		Lease:   msg.Lease,
		Height:  ctx.BlockHeight(),
	}

	// Checks that the element exists
	if !k.HasCrudValue(ctx, msg.Uuid, msg.Key) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Key))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetCrudValueOwner(ctx, msg.Uuid, msg.Key) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetCrudValue(ctx, CrudValue)

	return &types.MsgUpdateCrudValueResponse{}, nil
}

func (k msgServer) DeleteCrudValue(goCtx context.Context, msg *types.MsgDeleteCrudValue) (*types.MsgDeleteCrudValueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasCrudValue(ctx, msg.Uuid, msg.Key) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Key))
	}
	if msg.Creator != k.GetCrudValueOwner(ctx, msg.Uuid, msg.Key) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveCrudValue(ctx, msg.Uuid, msg.Key)

	return &types.MsgDeleteCrudValueResponse{}, nil
}