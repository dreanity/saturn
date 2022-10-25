package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// GiveawaysCountByOrganizerKeyPrefix is the prefix to retrieve all GiveawaysCountByOrganizer
	GiveawaysCountByOrganizerKeyPrefix = "GiveawaysCountByOrganizer/value/"
)

// GiveawaysCountByOrganizerKey returns the store key to retrieve a GiveawaysCountByOrganizer from the index fields
func GiveawaysCountByOrganizerKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
