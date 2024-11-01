package uuidv7

import (
	"crypto/rand"
	"encoding/binary"
	"time"
)

// GenerateUUIDv7 generates a UUIDv7 based on a timestamp. If no timestamp is provided,
// it uses the current time.
func GenerateUUIDv7(timestamp ...time.Time) (string, error) {
	var ts time.Time
	if len(timestamp) > 0 {
		ts = timestamp[0]
	} else {
		ts = time.Now()
	}

	// Get milliseconds since Unix epoch.
	ms := ts.UnixMilli()

	// Allocate 16 bytes for UUID.
	uuid := make([]byte, 16)

	// Set the timestamp part (48 bits, or 6 bytes)
	binary.BigEndian.PutUint64(uuid[:8], uint64(ms)<<16)

	// Set the version bits (UUID v7 = 0111) in the 7th byte (0x70 = 0111 0000).
	uuid[6] = (uuid[6] & 0x0F) | 0x70

	// Set the variant bits in the 9th byte (0x80 = 1000 0000).
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	// Fill the remaining 10 bytes with random values.
	if _, err := rand.Read(uuid[8:]); err != nil {
		return "", err
	}

	// Manually construct the UUID string
	uuidStr := make([]byte, 36)
	hex := "0123456789abcdef"

	for i, j := 0, 0; i < 16; i++ {
		if i == 4 || i == 6 || i == 8 || i == 10 {
			uuidStr[j] = '-'
			j++
		}

		uuidStr[j] = hex[uuid[i]>>4]
		uuidStr[j+1] = hex[uuid[i]&0x0F]
		j += 2
	}

	return string(uuidStr), nil
}
