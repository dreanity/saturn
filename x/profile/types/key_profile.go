package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProfileKeyPrefix is the prefix to retrieve all Profile
	ProfileKeyPrefix = "Profile/value/"
)

// ProfileKey returns the store key to retrieve a Profile from the index fields
func ProfileKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
