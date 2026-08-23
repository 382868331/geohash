package geohash

import "testing"

func TestTask013DecodeRangeOffset(t *testing.T) {
	if got := decodeRange(0, 90); got != -90 {
		t.Fatalf("latitude minimum=%v", got)
	}
	if got := decodeRange(1<<31, 90); got != 0 {
		t.Fatalf("latitude midpoint=%v", got)
	}
	if got := decodeRange(0, 180); got != -180 {
		t.Fatalf("longitude minimum=%v", got)
	}
	if got := decodeRange(1<<31, 180); got != 0 {
		t.Fatalf("longitude midpoint=%v", got)
	}
}
