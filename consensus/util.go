package consensus

import (
	"encoding/binary"
)

// heightToKey convert block height to byte array
func heightToKey(height uint32) []byte {
	key := make([]byte, 4)
	binary.LittleEndian.PutUint32(key, height)
	return key
}
