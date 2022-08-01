package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// GiveawayKeyPrefix is the prefix to retrieve all Giveaway
	GiveawayKeyPrefix = "Giveaway/value/"
)

// GiveawayKey returns the store key to retrieve a Giveaway from the index fields
func GiveawayKey(
	index uint64,
) []byte {
	var key []byte

	indexBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(indexBytes, index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
