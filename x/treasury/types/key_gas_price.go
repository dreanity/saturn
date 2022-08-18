package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// GasPriceKeyPrefix is the prefix to retrieve all GasPrice
	GasPriceKeyPrefix = "GasPrice/value/"
)

// GasPriceKey returns the store key to retrieve a GasPrice from the index fields
func GasPriceKey(
	currency string,
) []byte {
	var key []byte

	currencyBytes := []byte(currency)
	key = append(key, currencyBytes...)
	key = append(key, []byte("/")...)

	return key
}
