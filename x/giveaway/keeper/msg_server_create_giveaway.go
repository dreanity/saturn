package keeper

import (
	"context"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/giveaway/types"
)

func (k msgServer) CreateGiveaway(goCtx context.Context, msg *types.MsgCreateGiveaway) (*types.MsgCreateGiveawayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.profileKeeper.GetProfile(ctx, msg.Creator)

	if !found {
		return nil, types.ErrNotFoundCreatorProfile
	}

	if msg.Duration < 0 { // TODO: Поправить время на 3600
		return nil, types.ErrDurationToLow
	}

	if len(msg.Name) == 0 {
		return nil, types.ErrNameIsShort
	}

	if len(msg.Prizes) == 0 {
		return nil, types.ErrPrizesMustBeMoreThanOne
	}

	for _, prize := range msg.Prizes {
		if prize.Amount.LTE(sdk.ZeroInt()) {
			return nil, types.ErrPrizeAmountMustBeMoreThanZero
		}

		if prize.Scale > 30 {
			return nil, types.ErrPrizeScaleToBig
		}

		if len(prize.Currency) == 0 {
			return nil, types.ErrPrizeCurrencyMustBeMoreThanZero
		}
	}

	giveawayCount := k.GetGiveawayCount(ctx)
	giveawayCount.Value += 1

	genTime := k.gentimeKeeper.GetTime(ctx)
	height := sdk.NewDecFromBigInt(big.NewInt(ctx.BlockHeight()))
	blockTime := ctx.BlockTime()
	timeDiff := sdk.NewDecFromBigInt(
		big.NewInt(blockTime.UTC().Unix() - genTime.Value))

	avgBlockTime := timeDiff.Quo(height)

	numberOfBlocksToComplete := sdk.NewDecFromBigInt(
		big.NewInt(msg.Duration)).Quo(avgBlockTime).Ceil()

	complitionHeight := height.Add(numberOfBlocksToComplete).TruncateInt().Int64()

	giveaway := types.Giveaway{
		Index:                giveawayCount.Value,
		Creator:              msg.Creator,
		Duration:             msg.Duration,
		CreatedAt:            ctx.BlockTime().UTC().Unix(),
		Name:                 msg.Name,
		Prizes:               msg.Prizes,
		CompletionHeight:     complitionHeight,
		Status:               types.GiveawayStatus_TICKETS_REGISTRATION,
		WinningTicketNumbers: []uint32{},
	}

	k.SetGiveawayCount(ctx, giveawayCount)
	k.SetGiveaway(ctx, giveaway)

	giveawaysByHeight, found := k.GetGiveawayByHeight(ctx, giveaway.CompletionHeight)
	if !found {
		giveawaysByHeight.Height = giveaway.CompletionHeight
		giveawaysByHeight.Indexes = append(giveawaysByHeight.Indexes, giveaway.Index)
	} else {
		giveawaysByHeight.Indexes = append(giveawaysByHeight.Indexes, giveaway.Index)
	}
	k.SetGiveawayByHeight(ctx, giveawaysByHeight)

	ticketCount := types.TicketCount{
		GiveawayId: giveaway.Index,
		Count:      0,
	}
	k.SetTicketCount(ctx, ticketCount)

	ctx.EventManager().EmitTypedEvent(&types.GiveawayCreated{
		Index:                giveaway.Index,
		Creator:              giveaway.Creator,
		Duration:             giveaway.Duration,
		CreatedAt:            giveaway.CreatedAt,
		Name:                 giveaway.Name,
		Prizes:               giveaway.Prizes,
		CompletionHeight:     giveaway.CompletionHeight,
		Status:               giveaway.Status,
		WinningTicketNumbers: giveaway.WinningTicketNumbers,
	})

	return &types.MsgCreateGiveawayResponse{}, nil
}
