package uuid_test

import (
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/bentranter/uuid"
	satori "github.com/satori/go.uuid"
)

var buf [16]byte

func Example() {
	id := uuid.NewV4()
	fmt.Printf("%s\n", id)
}

func TestNewV1(t *testing.T) {
	x := uuid.NewV1().String()
	y := uuid.NewV1().String()

	if x == y {
		t.Fatalf("Values %s and %s\n are the same\n", x, y)
	}
}

func TestNewV4(t *testing.T) {
	x := uuid.NewV4().String()
	y := uuid.NewV4().String()

	if x == y {
		t.Fatalf("Values %s and %s\n are the same\n", x, y)
	}
}

func BenchmarkNewV1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		uuid.NewV1()
	}
}

func BenchmarkNewV4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		uuid.NewV4()
	}
}

func BenchmarkNewSatoriV1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		satori.NewV1()
	}
}

func BenchmarkNewSatoriV4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		satori.NewV4()
	}
}

func BenchmarkRandAlloc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		rand.Read(buf[:])
	}
}
