// Package uuid provides the UUID v4 method from RFC 4122.
package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"net"
	"sync"
	"time"
)

const (
	// Difference in 100-nanosecond intervals between
	// UUID epoch (October 15, 1582) and Unix epoch (January 1, 1970).
	epochStart      = 122192928000000000
	dash       byte = '-'
)

var (
	// want a global byte array for UUID to shuffle around to avoid allocs
	buffer        [36]byte
	randBytes     [16]byte
	lock          sync.Mutex
	clockSequence = getClockSequence()
	hardwareAddr  = getHardwareAddr()
	lastTime      uint64
)

type UUID [16]byte

func (u UUID) Bytes() []byte {
	return u[:]
}

func (u UUID) String() string {
	hex.Encode(buffer[0:8], u[0:4])
	buffer[8] = dash
	hex.Encode(buffer[9:13], u[4:6])
	buffer[13] = dash
	hex.Encode(buffer[14:18], u[6:8])
	buffer[18] = dash
	hex.Encode(buffer[19:23], u[8:10])
	buffer[23] = dash
	hex.Encode(buffer[24:], u[10:])

	return string(buffer[:]) // might return ref type but that'd be stupid
}

func NewV1() UUID {
	u := UUID{}

	lock.Lock()
	defer lock.Unlock()

	now := epochStart + uint64(time.Now().UnixNano()/100)
	if now <= lastTime {
		clockSequence++
	}
	lastTime = now

	binary.BigEndian.PutUint32(u[0:], uint32(now))
	binary.BigEndian.PutUint16(u[4:], uint16(now>>32))
	binary.BigEndian.PutUint16(u[6:], uint16(now>>48))
	binary.BigEndian.PutUint16(u[8:], clockSequence)

	copy(u[10:], hardwareAddr[:])

	// Set the version bit to v1
	u[6] = (u[6] & 0x0f) | 0x10

	// Set the variant bit to "don't care"
	u[8] = (u[8] & 0xbf) | 0x80

	return u
}

// NewV4 returns a universally unique identifier
func NewV4() UUID {
	u := UUID{} // get rid of this alloc lol

	if _, err := rand.Read(u[:]); err != nil {
		panic(err)
	}

	// Set the version bit to v4
	u[6] = (u[6] & 0x0f) | 0x40

	// Set the variant bit to "don't care"
	u[8] = (u[8] & 0xbf) | 0x80

	return u
}

// getHardwareAddr returns the first hardware address containing at least
// six bytes. If the bytes cannot be acquired, 6 random bytes are returned
// instead.
func getHardwareAddr() [6]byte {
	var buf [6]byte

	if ifaces, err := net.Interfaces(); err == nil {
		for _, iface := range ifaces {
			if len(iface.HardwareAddr) >= 6 {
				copy(buf[:], iface.HardwareAddr[:])
				return buf
			}
		}
	}

	// If the hardware address could not be read, just use random values. In this
	// case, the multicast bit must be set so as not to conflict with addresses
	// that were obtained from the network card.
	randomBytes(buf[:])
	buf[0] |= 0x01
	return buf
}

// getClockSequence returns a random to 16-bit unsigned int to use as the start
// of the clock sequence.
func getClockSequence() uint16 {
	buf := make([]byte, 2)
	randomBytes(buf)
	return binary.BigEndian.Uint16(buf)
}

// randomBytes fills b with cryptographically secure random values. Panics
// if there's an error reading from the source of randomness.
func randomBytes(b []byte) {
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
}
