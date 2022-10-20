package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/giveaway/types"
)

func (k msgServer) IssueTicket(goCtx context.Context, msg *types.MsgIssueTicket) (*types.MsgIssueTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if len(strings.Trim(msg.ParticipantId, " ")) < 1 {
		return nil, types.ErrInvalidParticipantId
	}

	if len(strings.Trim(msg.ParticipantName, " ")) < 1 {
		return nil, types.ErrInvalidParticipantName
	}

	giveaway, found := k.GetGiveaway(ctx, msg.GiveawayId)
	if !found {
		return nil, types.ErrIssueTicketForNonExistentGiveaway
	}

	if giveaway.Creator != msg.Creator {
		return nil, types.ErrYouAreNotTheCreator
	}

	if giveaway.Status != types.GiveawayStatus_TICKETS_REGISTRATION {
		return nil, types.ErrTicketRegistrationClosed
	}

	ticketCount, found := k.GetTicketCount(ctx, msg.GiveawayId)
	if !found {
		return nil, types.ErrIssueTicketForNonExistentGiveaway
	}

	ticket := types.Ticket{
		GiveawayId:      msg.GiveawayId,
		Index:           ticketCount.Count,
		ParticipantId:   msg.ParticipantId,
		ParticipantName: msg.ParticipantName,
	}
	ticketCount.Count += 1

	k.SetTicket(ctx, ticket)
	k.SetTicketCount(ctx, ticketCount)

	ctx.EventManager().EmitTypedEvent(&types.TicketCreated{
		GiveawayId:      ticket.GiveawayId,
		Index:           ticket.Index,
		ParticipantId:   ticket.ParticipantId,
		ParticipantName: ticket.ParticipantName,
	})

	return &types.MsgIssueTicketResponse{}, nil
}
