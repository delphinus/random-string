// Package randomString is written with referring this entry:
// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
package randomString

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

// LetterBytes means letters to be used in random string. It contains all of
// alphabets and numbers, but lacks undistinguishable letters, 1, l, I, O, 0.
//
// original: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const LetterBytes = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"

var m sync.Mutex
var src = rand.NewSource(time.Now().UnixNano())                    // source for creating random numbers
var letterIdxBits = uint(math.Log2(float64(len(LetterBytes)))) + 1 // bits to represent a letter index
var letterIdxMask = int64(1<<letterIdxBits - 1)                    // All 1-bits, as many as letterIdxBits
var letterIdxMax = 63 / letterIdxBits                              // # of letter indices fitting in 63 bits

// Generate generates random string
func Generate(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax characters!
	m.Lock()
	cache := src.Int63()
	m.Unlock()
	for i, remain := n-1, letterIdxMax; i >= 0; {
		if remain == 0 {
			m.Lock()
			cache = src.Int63()
			m.Unlock()
			remain = letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(LetterBytes) {
			b[i] = LetterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
