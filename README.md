# RandomString -- Library to make passwords without misleading letters

This package makes random passwords with fast logic described in [this link][].

[this link]: http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang

The original logic is NOT goroutine safe because it shares `rand.Source` and the `Int63()` calls conflict. The logic in this package is using `sync.Mutex` to solve this.

```go
package main

import (
  "fmt"

  "github.com/delphinus/random-string"
)

func main() {
  pass := randomString.Generate(8);
  fmt.Printf("simple password: %s\n", pass)
	// simple password: JEVGqkiW
}
```
