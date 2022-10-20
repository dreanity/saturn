package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/profile/types"
)

func (k msgServer) UpdateProfile(goCtx context.Context, msg *types.MsgUpdateProfile) (*types.MsgUpdateProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	emgr := ctx.EventManager()

	if len(msg.Name) == 0 {
		return nil, types.ErrNameIsShort
	}

	profile, hasProfile := k.GetProfile(ctx, msg.Creator)
	nameData, hasNameData := k.GetNameRegistry(ctx, msg.Name)

	if !hasProfile {
		if hasNameData {
			return nil, types.ErrNameAlreadyUsed
		}

		nameData = types.NameRegistry{
			Name:    msg.Name,
			Address: msg.Creator,
		}

		profile = types.Profile{
			Address:   msg.Creator,
			AvatarUrl: msg.AvatarUrl,
			BannerUrl: msg.BannerUrl,
			Name:      msg.Name,
		}
	} else {
		if hasNameData && nameData.Address != msg.Creator {
			return nil, types.ErrNameAlreadyUsed
		}

		if !hasNameData {
			k.RemoveNameRegistry(ctx, profile.Name)
		}

		nameData.Name = msg.Name
		nameData.Address = msg.Creator
		profile.AvatarUrl = msg.AvatarUrl
		profile.BannerUrl = msg.BannerUrl
		profile.Name = msg.Name
	}

	k.SetNameRegistry(ctx, nameData)
	k.SetProfile(ctx, profile)
	emgr.EmitTypedEvent(&types.ProfileUpdated{
		Address:   msg.Creator,
		AvatarUrl: msg.AvatarUrl,
		BannerUrl: msg.BannerUrl,
		Name:      msg.Name,
	})

	return &types.MsgUpdateProfileResponse{}, nil
}
