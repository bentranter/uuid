# UUID

Generate UUID v4s without any allocations.

### Usage

```go
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
```

### Benchmark

```
go test -bench=. -benchmem
BenchmarkGenerate-4   	 1000000	      2083 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/bentranter/uuid	2.119s
```

### License

Apache 2. See the LICENSE file for a copy.
