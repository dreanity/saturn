package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/profile/types"
)

func (k msgServer) UpdateProfile(goCtx context.Context, msg *types.MsgUpdateProfile) (*types.MsgUpdateProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	profile, hasProfile := k.GetProfile(ctx, msg.Creator)
	if !hasProfile {
		profile = types.Profile{
			Address:   msg.Creator,
			AvatarUrl: msg.AvatarUrl,
			BannerUrl: msg.BannerUrl,
			Name:      msg.Name,
		}
	} else {
		profile.AvatarUrl = msg.AvatarUrl
		profile.BannerUrl = msg.BannerUrl
		profile.Name = msg.Name
	}

	k.SetProfile(ctx, profile)

	return &types.MsgUpdateProfileResponse{}, nil
}
