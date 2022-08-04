package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/giveaway/types"
)

func (k msgServer) IssueTicket(goCtx context.Context, msg *types.MsgIssueTicket) (*types.MsgIssueTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ticketCount, found := k.GetTicketCount(ctx, msg.GiveawayId)
	if !found {
		return nil, types.ErrIssueTicketForNonExistentGiveaway
	}
	ticketCount.Count += 1

	ticket := types.Ticket{
		Index:           ticketCount.Count,
		GiveawayId:      msg.GiveawayId,
		ParticipantId:   msg.ParticipantId,
		ParticipantName: msg.ParticipantName,
	}

	k.SetTicket(ctx, ticket)
	k.SetTicketCount(ctx, ticketCount)

	return &types.MsgIssueTicketResponse{}, nil
}
