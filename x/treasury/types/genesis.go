package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Treasurer: Treasurer{
			Address: "saturn1gg2tcp9c6sv2w3uuyy23ncy90re4tf6d97n7ev",
		},
		GasPriceList: []GasPrice{},
		GasBidList:   []GasBid{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in gasPrice
	gasPriceIndexMap := make(map[string]struct{})

	for _, elem := range gs.GasPriceList {
		index := string(GasPriceKey(elem.Chain, elem.TokenAddress))
		if _, ok := gasPriceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for gasPrice")
		}
		gasPriceIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in gasBid
	gasBidIndexMap := make(map[string]struct{})

	for _, elem := range gs.GasBidList {
		index := string(GasBidKey(elem.Chain))
		if _, ok := gasBidIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for gasBid")
		}
		gasBidIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
