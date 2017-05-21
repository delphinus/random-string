package randomString

// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-testLength-in-golang

import (
	"math"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func BenchmarkRandomString(b *testing.B) {
	length := 256
	letterBytes := "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	letterRunes := []rune(letterBytes)
	m := sync.Mutex{}
	src := rand.NewSource(time.Now().UnixNano())
	bits := uint(math.Log2(float64(len(letterBytes)))) + 1
	mask := int64(1<<bits - 1)
	max := 63 / bits

	randStringRunes := func(n int) string {
		by := make([]rune, n)
		for i := range by {
			by[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
		return string(by)
	}

	randStringBytes := func(n int) string {
		by := make([]byte, n)
		for i := range by {
			by[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
		return string(by)
	}

	randStringRemainder := func(n int) string {
		by := make([]byte, n)
		for i := range by {
			by[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
		}
		return string(by)
	}

	randStringMask := func(n int) string {
		by := make([]byte, n)
		for i := 0; i < n; {
			if idx := int(rand.Int63() & mask); idx < len(letterBytes) {
				by[i] = letterBytes[idx]
				i++
			}
		}
		return string(by)
	}

	randStringMaskImproved := func(n int) string {
		by := make([]byte, n)
		for i, cache, remain := n-1, rand.Int63(), max; i >= 0; {
			if remain == 0 {
				cache, remain = rand.Int63(), max
			}
			if idx := int(cache & mask); idx < len(letterBytes) {
				by[i] = letterBytes[idx]
				i--
			}
			cache >>= bits
			remain--
		}
		return string(by)
	}

	randStringSource := func(n int) string {
		by := make([]byte, n)
		m.Lock()
		cache := src.Int63()
		m.Unlock()
		for i, remain := n-1, max; i >= 0; {
			if remain == 0 {
				m.Lock()
				cache = src.Int63()
				m.Unlock()
				remain = max
			}
			if idx := int(cache & mask); idx < len(letterBytes) {
				by[i] = letterBytes[idx]
				i--
			}
			cache >>= bits
			remain--
		}
		return string(by)
	}

	b.Run("Runes", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = randStringRunes(length)
			}
		})
	})

	b.Run("Bytes", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = randStringBytes(length)
			}
		})
	})

	b.Run("Remainder", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = randStringRemainder(length)
			}
		})
	})

	b.Run("Mask", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = randStringMask(length)
			}
		})
	})

	b.Run("MaskImproved", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = randStringMaskImproved(length)
			}
		})
	})

	b.Run("MaskSource", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = randStringSource(length)
			}
		})
	})
}
