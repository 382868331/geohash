package geohash

import "testing"

func TestTask007PrecisionControlsExactLength(t *testing.T) {
	for _, chars := range []uint{1, 5, 12} {
		got := EncodeWithPrecision(42.6, -5.6, chars)
		if uint(len(got)) != chars {
			t.Fatalf("precision %d produced %q with length %d", chars, got, len(got))
		}
	}
}
