package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UnprovenRandomnessKeyPrefix is the prefix to retrieve all UnprovenRandomness
	UnprovenRandomnessKeyPrefix = "UnprovenRandomness/value/"
)

// UnprovenRandomnessKey returns the store key to retrieve a UnprovenRandomness from the index fields
func UnprovenRandomnessKey(
	round uint64,
) []byte {
	var key []byte

	roundBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(roundBytes, round)
	key = append(key, roundBytes...)
	key = append(key, []byte("/")...)

	return key
}
