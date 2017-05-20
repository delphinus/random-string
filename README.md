# RandomString -- Library to make passwords without misleading letters

[![CircleCI](https://circleci.com/gh/delphinus/random-string.svg?style=svg)](https://circleci.com/gh/delphinus/random-string)
[![Coverage Status](https://coveralls.io/repos/github/delphinus/random-string/badge.svg)](https://coveralls.io/github/delphinus/random-string)

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

## Benchmark

```go
go version
go version go1.8.1 darwin/amd64
go test -bench . -benchmem
BenchmarkRandomString/Runes-8              50000             35584 ns/op            1312 B/op          2 allocs/op
BenchmarkRandomString/Bytes-8              50000             35547 ns/op             512 B/op          2 allocs/op
BenchmarkRandomString/Remainder-8          50000             34095 ns/op             512 B/op          2 allocs/op
BenchmarkRandomString/Mask-8               50000             35362 ns/op             512 B/op          2 allocs/op
BenchmarkRandomString/MaskImproved-8      300000              4710 ns/op             512 B/op          2 allocs/op
BenchmarkRandomString/MaskSource-8        500000              3873 ns/op             512 B/op          2 allocs/op
PASS
ok      github.com/delphinus/random-string      11.885s
```
