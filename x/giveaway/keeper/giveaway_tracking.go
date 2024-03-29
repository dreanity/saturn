package keeper

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"

	stdcrypto "crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/crypto"
	"github.com/dreanity/saturn/x/giveaway/types"
)

func (k Keeper) TrackGiveawayByHeight(ctx sdk.Context) {
	height := ctx.BlockHeight()
	giveawaysByHeight, found := k.GetGiveawayByHeight(ctx, height)
	if found {
		emgr := ctx.EventManager()

		blockTime := ctx.BlockTime()
		randomnessRound := k.randomnessKeeper.ComputeRandomnessRoundForTime(ctx, uint64(blockTime.UTC().Unix()))

		giveawayByRandomness, found := k.GetGiveawayByRandomness(ctx, randomnessRound)
		if !found {
			giveawayByRandomness.Round = randomnessRound
		}

		unprovenRandomnessRequired := false
		for _, index := range giveawaysByHeight.Indexes {
			giveaway, found := k.GetGiveaway(ctx, index)
			if !found {
				panic("Giveaway not found by giveawayByGiveaway indexes")
			}

			ticketCount, found := k.GetTicketCount(ctx, index)
			if !found {
				msg := fmt.Sprintf("Ticket count for giveway #%d not found", index)
				panic(msg)
			}

			if ticketCount.Count == 0 || ticketCount.Count < uint32(len(giveaway.Prizes)) {
				giveaway.Status = types.GiveawayStatus_CANCELLED_INSUF_TICKETS
				k.SetGiveaway(ctx, giveaway)
				emgr.EmitTypedEvents(&types.GiveawayCancelledInsufTickets{
					GiveawayId: giveaway.Index,
				})
				continue
			}

			giveawayByRandomness.Indexes = append(giveawayByRandomness.Indexes, index)

			giveaway.Status = types.GiveawayStatus_WINNERS_DETERMINATION
			k.SetGiveaway(ctx, giveaway)
			unprovenRandomnessRequired = true

			emgr.EmitTypedEvents(&types.GiveawayWinnersDeterminationBegun{
				GiveawayId:      giveaway.Index,
				RandomnessRound: randomnessRound,
			})
		}

		if unprovenRandomnessRequired {
			unprovenRandomness, found := k.randomnessKeeper.GetUnprovenRandomness(ctx, randomnessRound)

			if !found {
				unprovenRandomness.Round = randomnessRound

				roundTime := k.randomnessKeeper.ComputeTimeForRandomnessRound(ctx, randomnessRound)
				unprovenRandomness.RoundTime = roundTime

				k.randomnessKeeper.SetUnprovenRandomnessWithEvent(ctx, unprovenRandomness)
			}
		}

		k.SetGiveawayByRandomness(ctx, giveawayByRandomness)
		k.RemoveGiveawayByHeight(ctx, height)
	}
}

func (k Keeper) TrackGiveawayByRandomness(ctx sdk.Context) {
	giveawaysByRandomness := k.GetAllGiveawayByRandomness(ctx)
	emgr := ctx.EventManager()

	for _, giveawayByRandomness := range giveawaysByRandomness {
		round := giveawayByRandomness.Round
		provenRandomness, found := k.randomnessKeeper.GetProvenRandomness(ctx, round)

		if found {
			for _, giveawayIndex := range giveawayByRandomness.Indexes {
				giveaway, found := k.GetGiveaway(ctx, giveawayIndex)
				if !found {
					msg := fmt.Sprintf("Giveaway #%d not found", giveawayIndex)
					panic(msg)
				}

				ticketCount, found := k.GetTicketCount(ctx, giveawayIndex)
				if !found {
					msg := fmt.Sprintf("Ticket count for giveway #%d not found", giveawayIndex)
					panic(msg)
				}

				if ticketCount.Count == 0 || ticketCount.Count < uint32(len(giveaway.Prizes)) {
					giveaway.Status = types.GiveawayStatus_CANCELLED_INSUF_TICKETS

					continue
				}

				drbgSeed, err := hex.DecodeString(provenRandomness.Randomness)
				if err != nil {
					panic("Randomness not hex decode")
				}
				drbgNonce := make([]byte, 4)
				binary.BigEndian.PutUint32(drbgNonce, giveawayIndex)

				drbgPersonalize := []byte(ctx.ChainID())

				winnersCount := len(giveaway.Prizes)
				winnersNumber := determineWinners(ctx, drbgSeed, drbgNonce, drbgPersonalize, ticketCount.Count, winnersCount)

				giveaway.WinningTicketNumbers = winnersNumber
				giveaway.Status = types.GiveawayStatus_WINNERS_DETERMINED

				k.SetGiveaway(ctx, giveaway)
				emgr.EmitTypedEvents(&types.GiveawayWinnersDetermined{
					GiveawayId:     giveaway.Index,
					WinnersNumbers: winnersNumber,
				})
			}

			k.RemoveGiveawayByRandomness(ctx, round)

		}
	}
}

func determineWinners(ctx sdk.Context, seed, nonce, personalize []byte, ticketsCount uint32, winnersCount int) []uint32 {
	drbg, err := crypto.NewHmacDrbg(stdcrypto.SHA256, seed, nonce, personalize)
	if err != nil {
		msg := fmt.Sprintf("error creating hmac_drbg: %s", err)
		panic(msg)
	}

	winnerNumbersToPlaces := make(map[uint32]uint32) // TICKET_NUMBER -> WINNING_PLACE

	for place := 0; place < winnersCount; place++ {
		winnerNumber := drbg.ReadUint32() % ticketsCount

		_, alreadyWinning := winnerNumbersToPlaces[winnerNumber]
		if alreadyWinning {
			leftWinnerNumber := winnerNumber
			rightWinnerNumber := winnerNumber

			for {
				if leftWinnerNumber != 0 {
					leftWinnerNumber -= 1
				}

				if rightWinnerNumber != ticketsCount-1 {
					rightWinnerNumber += 1
				}

				_, leftWinnerNumberAlreadyWinning := winnerNumbersToPlaces[leftWinnerNumber]
				_, rightWinnerNumberAlreadyWinning := winnerNumbersToPlaces[rightWinnerNumber]

				if leftWinnerNumberAlreadyWinning && rightWinnerNumberAlreadyWinning {
					continue
				}

				if !leftWinnerNumberAlreadyWinning && !rightWinnerNumberAlreadyWinning {
					if winnerNumber%2 == 0 {
						winnerNumber = leftWinnerNumber
					} else {
						winnerNumber = rightWinnerNumber
					}

					break
				}

				if !leftWinnerNumberAlreadyWinning {
					winnerNumber = leftWinnerNumber
					break
				}

				if !rightWinnerNumberAlreadyWinning {
					winnerNumber = rightWinnerNumber
					break
				}
			}
		}

		winnerNumbersToPlaces[winnerNumber] = uint32(place)
	}

	winnersNumber := make([]uint32, 0)
	for winnerNumber, _ := range winnerNumbersToPlaces {
		winnersNumber = append(winnersNumber, winnerNumber)
	}

	return winnersNumber
}
