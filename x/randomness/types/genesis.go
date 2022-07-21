package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ChainInfo: ChainInfo{
			PublicKey:   "868f005eb8e6e4ca0a47c8a77ceaa5309a47978a7c71bc5cce96366b5d7a569937c529eeda66c7293784a9402801af31",
			Period:      30,
			GenesisTime: 1595431050,
			Hash:        "8990e7a9aaed2ffed73dbd7092123d6f289930540d7651336225dc172e51b2ce",
		},
		UnprovenRendomnessList: []UnprovenRendomness{},
		ProvenRandomnessList:   []ProvenRandomness{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in unprovenRendomness
	unprovenRendomnessIndexMap := make(map[string]struct{})

	for _, elem := range gs.UnprovenRendomnessList {
		index := string(UnprovenRendomnessKey(elem.Index))
		if _, ok := unprovenRendomnessIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for unprovenRendomness")
		}
		unprovenRendomnessIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in provenRandomness
	provenRandomnessIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProvenRandomnessList {
		index := string(ProvenRandomnessKey(elem.Index))
		if _, ok := provenRandomnessIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for provenRandomness")
		}
		provenRandomnessIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
