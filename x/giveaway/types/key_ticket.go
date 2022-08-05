package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	// TicketKeyPrefix is the prefix to retrieve all Ticket
	TicketKeyPrefix = "Ticket/value/"
)

// TicketKey returns the store key to retrieve a Ticket from the index fields
func TicketKey(
	giveawayId uint32,
	index uint32,
) []byte {
	return []byte(fmt.Sprintf("%d/%d", giveawayId, index))

	// var key []byte

	// giveawayBytes := make([]byte, 4)
	// binary.BigEndian.PutUint32(giveawayBytes, giveawayId)

	// indexBytes := make([]byte, 4)
	// binary.BigEndian.PutUint32(indexBytes, index)

	// key = append(key, giveawayBytes...)
	// key = append(key, []byte("/")...)
	// key = append(key, indexBytes...)

	// return key
}
