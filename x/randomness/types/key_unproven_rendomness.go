package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UnprovenRandomnessKeyPrefix is the prefix to retrieve all UnprovenRandomness
	UnprovenRandomnessKeyPrefix = "UnprovenRandomness/value/"
)

// UnprovenRandomnessKey returns the store key to retrieve a UnprovenRandomness from the index fields
func UnprovenRandomnessKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
