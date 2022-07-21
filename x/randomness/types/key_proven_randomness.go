package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProvenRandomnessKeyPrefix is the prefix to retrieve all ProvenRandomness
	ProvenRandomnessKeyPrefix = "ProvenRandomness/value/"
)

// ProvenRandomnessKey returns the store key to retrieve a ProvenRandomness from the index fields
func ProvenRandomnessKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
