package main

import (
	"fmt"

	"github.com/bentranter/uuid"
)

func main() {
	x := uuid.V4()
	y := uuid.V4()

	fmt.Printf("%s %s\n", x, y)
}
