# RandomString -- Library to make passwords without misleading letters

This package makes random passwords with fast logic described in [this link][].

[this link]: http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang

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
