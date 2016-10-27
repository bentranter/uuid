# UUID

Generate UUID v4s without any allocations.

### Usage

_(Don't actually use this, use https://github.com/satori/go.uuid instead, it's much better)_

```go
package main

import (
	"fmt"

	"github.com/bentranter/uuid"
)

func main() {
	x := uuid.V4()
	y := uuid.V4()

	fmt.Println(string(x[:]))
	fmt.Println(string(y[:]))

	// Output (example):
	// ee416737-4a97-49aa-a3e9-b6fce0bb757a
	// 51033a45-4a43-4551-848b-e6f246541652
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
