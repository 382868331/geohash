package geohash

import "testing"

func TestTask014ContainsUsesLongitudeForBothBounds(t *testing.T) {
	box := Box{MinLat: 40, MaxLat: 50, MinLng: -10, MaxLng: 10}
	if !box.Contains(45, 0) {
		t.Fatal("center point should be contained")
	}
	if box.Contains(45, -20) {
		t.Fatal("longitude below the western edge must be rejected")
	}
}
