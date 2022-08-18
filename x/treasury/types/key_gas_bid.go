package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// GasBidKeyPrefix is the prefix to retrieve all GasBid
	GasBidKeyPrefix = "GasBid/value/"
)

// GasBidKey returns the store key to retrieve a GasBid from the index fields
func GasBidKey(
	fromChain string,
) []byte {
	var key []byte

	fromChainBytes := []byte(fromChain)
	key = append(key, fromChainBytes...)
	key = append(key, []byte("/")...)

	return key
}
