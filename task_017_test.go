package geohash

import "testing"

func TestTask017IntegerBoxContainsEncodedPoint(t *testing.T) {
	for _, p := range [][2]float64{{57.64911, 10.40744}, {-33.86, 151.21}} {
		h := EncodeIntWithPrecision(p[0], p[1], 30)
		b := BoundingBoxIntWithPrecision(h, 30)
		if !b.Contains(p[0], p[1]) {
			t.Fatalf("point=%v box=%#v", p, b)
		}
	}
	h := EncodeIntWithPrecision(0, 0, 20)
	if !BoundingBoxIntWithPrecision(h, 20).Contains(0, 0) {
		t.Fatal("origin missing")
	}
}
