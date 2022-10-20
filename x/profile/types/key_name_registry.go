package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NameRegistryKeyPrefix is the prefix to retrieve all NameRegistry
	NameRegistryKeyPrefix = "NameRegistry/value/"
)

// NameRegistryKey returns the store key to retrieve a NameRegistry from the index fields
func NameRegistryKey(
	name string,
) []byte {
	var key []byte

	nameBytes := []byte(name)
	key = append(key, nameBytes...)
	key = append(key, []byte("/")...)

	return key
}
