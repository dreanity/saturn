package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TicketCountKeyPrefix is the prefix to retrieve all TicketCount
	TicketCountKeyPrefix = "TicketCount/value/"
)

// TicketCountKey returns the store key to retrieve a TicketCount from the index fields
func TicketCountKey(
	giveawayId uint32,
) []byte {
	var key []byte

	giveawayIdBytes := make([]byte, 8)
	binary.BigEndian.PutUint32(giveawayIdBytes, giveawayId)
	key = append(key, giveawayIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
