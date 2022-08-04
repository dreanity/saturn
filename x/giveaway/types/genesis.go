package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		GiveawayList:             []Giveaway{},
		GiveawayCount:            GiveawayCount{Value: 0},
		GiveawayByHeightList:     []GiveawayByHeight{},
		GiveawayByRandomnessList: []GiveawayByRandomness{},
		TicketList:               []Ticket{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in giveaway
	giveawayIndexMap := make(map[string]struct{})

	for _, elem := range gs.GiveawayList {
		index := string(GiveawayKey(elem.Index))
		if _, ok := giveawayIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for giveaway")
		}
		giveawayIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in giveawayByHeight
	giveawayByHeightIndexMap := make(map[string]struct{})

	for _, elem := range gs.GiveawayByHeightList {
		index := string(GiveawayByHeightKey(elem.Height))
		if _, ok := giveawayByHeightIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for giveawayByHeight")
		}
		giveawayByHeightIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in giveawayByRandomness
	giveawayByRandomnessIndexMap := make(map[string]struct{})

	for _, elem := range gs.GiveawayByRandomnessList {
		index := string(GiveawayByRandomnessKey(elem.Round))
		if _, ok := giveawayByRandomnessIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for giveawayByRandomness")
		}
		giveawayByRandomnessIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in ticket
	ticketIndexMap := make(map[string]struct{})

	for _, elem := range gs.TicketList {
		index := string(TicketKey(elem.Index))
		if _, ok := ticketIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for ticket")
		}
		ticketIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
