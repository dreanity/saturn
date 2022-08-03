package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// GiveawayByHeightKeyPrefix is the prefix to retrieve all GiveawayByHeight
	GiveawayByHeightKeyPrefix = "GiveawayByHeight/value/"
)

// GiveawayByHeightKey returns the store key to retrieve a GiveawayByHeight from the index fields
func GiveawayByHeightKey(
	height int32,
) []byte {
	var key []byte

	heightBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(heightBytes, uint32(height))
	key = append(key, heightBytes...)
	key = append(key, []byte("/")...)

	return key
}
