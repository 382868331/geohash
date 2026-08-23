package geohash

import "testing"

func TestTask005ValidateChecksFirstByte(t *testing.T) {
	for _, h := range []string{"a123", "i9", "o"} {
		if err := Validate(h); err == nil {
			t.Fatalf("invalid hash %q accepted", h)
		}
	}
	if err := Validate("u4pr"); err != nil {
		t.Fatal(err)
	}
	if err := Validate(""); err != nil {
		t.Fatalf("empty hash rejected: %v", err)
	}
}
