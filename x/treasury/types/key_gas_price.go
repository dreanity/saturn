package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// GasPriceKeyPrefix is the prefix to retrieve all GasPrice
	GasPriceKeyPrefix = "GasPrice/value/"
)

// GasPriceKey returns the store key to retrieve a GasPrice from the index fields
func GasPriceKey(
	chain string,
	tokenAddress string,
) []byte {
	var key []byte

	chainBytes := []byte(chain)
	currencyBytes := []byte(chain)

	key = append(key, chainBytes...)
	key = append(key, []byte("/")...)
	key = append(key, currencyBytes...)
	key = append(key, []byte("/")...)

	return key
}
