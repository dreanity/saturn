package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProvenRandomnessKeyPrefix is the prefix to retrieve all ProvenRandomness
	ProvenRandomnessKeyPrefix = "ProvenRandomness/value/"
)

// ProvenRandomnessKey returns the store key to retrieve a ProvenRandomness from the index fields
func ProvenRandomnessKey(
	round uint64,
) []byte {
	var key []byte

	roundBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(roundBytes, round)
	key = append(key, roundBytes...)
	key = append(key, []byte("/")...)

	return key
}
