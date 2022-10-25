package types_test

import (
	"testing"

	"github.com/dreanity/saturn/x/giveaway/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				GiveawayList: []types.Giveaway{
					{
						Index: 0,
					},
					{
						Index: 1,
					},
				},
				GiveawayCount: types.GiveawayCount{
					Value: 22,
				},
				GiveawayByHeightList: []types.GiveawayByHeight{
					{
						Height: 0,
					},
					{
						Height: 1,
					},
				},
				GiveawayByRandomnessList: []types.GiveawayByRandomness{
					{
						Round: 0,
					},
					{
						Round: 1,
					},
				},
				TicketList: []types.Ticket{
					{
						Index: 0,
					},
					{
						Index: 1,
					},
				},
				TicketCountList: []types.TicketCount{
					{
						GiveawayId: 0,
					},
					{
						GiveawayId: 1,
					},
				},
				GiveawaysCountByOrganizerList: []types.GiveawaysCountByOrganizer{
					{
						Address: "0",
					},
					{
						Address: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated giveaway",
			genState: &types.GenesisState{
				GiveawayList: []types.Giveaway{
					{
						Index: 0,
					},
					{
						Index: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated giveawayByHeight",
			genState: &types.GenesisState{
				GiveawayByHeightList: []types.GiveawayByHeight{
					{
						Height: 0,
					},
					{
						Height: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated giveawayByRandomness",
			genState: &types.GenesisState{
				GiveawayByRandomnessList: []types.GiveawayByRandomness{
					{
						Round: 0,
					},
					{
						Round: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated ticket",
			genState: &types.GenesisState{
				TicketList: []types.Ticket{
					{
						Index: 0,
					},
					{
						Index: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated ticketCount",
			genState: &types.GenesisState{
				TicketCountList: []types.TicketCount{
					{
						GiveawayId: 0,
					},
					{
						GiveawayId: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated giveawaysCountByOrganizer",
			genState: &types.GenesisState{
				GiveawaysCountByOrganizerList: []types.GiveawaysCountByOrganizer{
					{
						Address: "0",
					},
					{
						Address: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
