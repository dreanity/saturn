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

	if giveaway.Status != types.GiveawayStatus_TICKETS_REGISTRATION {
		return nil, types.ErrTicketRegistrationClosed
	}

	ticketCount, found := k.GetTicketCount(ctx, msg.GiveawayId)
	if !found {
		return nil, types.ErrIssueTicketForNonExistentGiveaway
	}

	ticket := types.Ticket{
		Index:           ticketCount.Count,
		GiveawayId:      msg.GiveawayId,
		ParticipantId:   msg.ParticipantId,
		ParticipantName: msg.ParticipantName,
	}
	ticketCount.Count += 1

	k.SetTicket(ctx, ticket)
	k.SetTicketCount(ctx, ticketCount)

	return &types.MsgIssueTicketResponse{}, nil
}
