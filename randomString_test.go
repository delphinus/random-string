package randomString

import "testing"

func TestLength(t *testing.T) {
	length := 16
	p := Generate(length)
	if len(p) != length {
		t.Fatalf("length differs: %d != %d", p, length)
	}
}
