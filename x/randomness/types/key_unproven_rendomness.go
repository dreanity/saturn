package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UnprovenRendomnessKeyPrefix is the prefix to retrieve all UnprovenRendomness
	UnprovenRendomnessKeyPrefix = "UnprovenRendomness/value/"
)

// UnprovenRendomnessKey returns the store key to retrieve a UnprovenRendomness from the index fields
func UnprovenRendomnessKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
