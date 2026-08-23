package geohash

import "testing"

func TestTask001EncodingValidity(t *testing.T) {
	for _, b := range []byte{'0', 'b', 'z'} {
		if !base32encoding.ValidByte(b) {
			t.Fatalf("%q should be valid", b)
		}
	}
	for _, b := range []byte{'a', 'i', 'o'} {
		if base32encoding.ValidByte(b) {
			t.Fatalf("%q should be invalid", b)
		}
	}
	if base32encoding.ValidByte(0xff) {
		t.Fatal("0xff should be invalid")
	}
}
