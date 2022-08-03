package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// GiveawayByRandomnessKeyPrefix is the prefix to retrieve all GiveawayByRandomness
	GiveawayByRandomnessKeyPrefix = "GiveawayByRandomness/value/"
)

// GiveawayByRandomnessKey returns the store key to retrieve a GiveawayByRandomness from the index fields
func GiveawayByRandomnessKey(
	round uint64,
) []byte {
	var key []byte

	roundBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(roundBytes, round)
	key = append(key, roundBytes...)
	key = append(key, []byte("/")...)

	return key
}
