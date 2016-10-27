// Package uuid provides the UUID v4 method from RFC 4122.
package uuid

import (
	"crypto/rand"
	"encoding/hex"
)

const dash byte = '-'

var (
	buffer    [36]byte
	randBytes [16]byte
)

// V4 returns a universally unique identifier
func V4() [36]byte {
	if _, err := rand.Read(randBytes[:]); err != nil {
		panic(err)
	}

	randBytes[6] = (randBytes[6] & 0x0f) | 0x40
	randBytes[8] = (randBytes[8] & 0xbf) | 0x80

	hex.Encode(buffer[0:8], randBytes[0:4])
	buffer[8] = dash
	hex.Encode(buffer[9:13], randBytes[4:6])
	buffer[13] = dash
	hex.Encode(buffer[14:18], randBytes[6:8])
	buffer[18] = dash
	hex.Encode(buffer[19:23], randBytes[8:10])
	buffer[23] = dash
	hex.Encode(buffer[24:], randBytes[10:])

	return buffer
}
