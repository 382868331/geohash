package geohash

import "testing"

func TestTask009StringPrecisionBits(t *testing.T) {
	for _, h := range []string{"u", "u4", "u4pruy"} {
		_, bits := ConvertStringToInt(h)
		if bits != uint(5*len(h)) {
			t.Fatalf("%q bits=%d", h, bits)
		}
	}
	_, bits := ConvertStringToInt("")
	if bits != 0 {
		t.Fatalf("empty bits=%d", bits)
	}
}
