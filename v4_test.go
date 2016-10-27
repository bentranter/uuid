package uuid_test

import (
	"fmt"
	"testing"

	"github.com/bentranter/uuid"
)

func Example() {
	id := uuid.V4()
	fmt.Printf("%s\n", id)
}

func TestV4(t *testing.T) {
	x := uuid.V4()
	y := uuid.V4()

	if x == y {
		t.Fatalf("UUIDs matched %s %s\n", x, y)
	}
}

func BenchmarkV4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		uuid.V4()
	}
}
