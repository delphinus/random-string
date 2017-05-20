# RandomString -- Library to make passwords without misleading letters

[![CircleCI](https://circleci.com/gh/delphinus/random-string.svg?style=svg)](https://circleci.com/gh/delphinus/random-string)
[![Coverage Status](https://coveralls.io/repos/github/delphinus/random-string/badge.svg)](https://coveralls.io/github/delphinus/random-string)

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
