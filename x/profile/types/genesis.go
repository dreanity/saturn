package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ProfileList:      []Profile{},
		NameRegistryList: []NameRegistry{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in profile
	profileIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProfileList {
		index := string(ProfileKey(elem.Address))
		if _, ok := profileIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for profile")
		}
		profileIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in nameRegistry
	nameRegistryIndexMap := make(map[string]struct{})

	for _, elem := range gs.NameRegistryList {
		index := string(NameRegistryKey(elem.Name))
		if _, ok := nameRegistryIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for nameRegistry")
		}
		nameRegistryIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
