package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/giveaway/types"
)

func (k Keeper) TrackGiveawayByHeight(ctx sdk.Context) {
	height := ctx.BlockHeight()
	giveawaysByHeight, found := k.GetGiveawayByHeight(ctx, height)
	if found {
		blockTime := ctx.BlockTime()
		randomnessRound := k.randomnessKeeper.ComputeRandomnessRoundForTime(ctx, uint64(blockTime.UTC().Unix()))

		unprovenRandomness, found := k.randomnessKeeper.GetUnprovenRandomness(ctx, randomnessRound)

		if !found {
			unprovenRandomness.Round = randomnessRound
			k.randomnessKeeper.SetUnprovenRandomness(ctx, unprovenRandomness)
		}

		giveawayByRandomness, found := k.GetGiveawayByRandomness(ctx, randomnessRound)
		if !found {
			giveawayByRandomness.Round = randomnessRound
		}
		giveawayByRandomness.Indexes = append(giveawayByRandomness.Indexes, giveawaysByHeight.Indexes...)
		k.SetGiveawayByRandomness(ctx, giveawayByRandomness)

		for _, index := range giveawaysByHeight.Indexes {
			giveaway, found := k.GetGiveaway(ctx, index)
			if !found {
				panic("Giveaway not found by giveawayByGiveaway indexes")
			}
			giveaway.Status = types.Giveaway_WINNERS_DETERMINATION
			k.SetGiveaway(ctx, giveaway)
		}

		k.RemoveGiveawayByHeight(ctx, height)
	}
}

func (k Keeper) TrackGiveawayByRandomness(ctx sdk.Context) {
	giveawaysByRandomness := k.GetAllGiveawayByRandomness(ctx)

	for _, giveawayByRandomness := range giveawaysByRandomness {
		round := giveawayByRandomness.Round
		_, found := k.randomnessKeeper.GetProvenRandomness(ctx, round)
		if found {
			//

			k.RemoveGiveawayByRandomness(ctx, round)
		}
	}
}
