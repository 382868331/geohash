package geohash

import "testing"

func TestTask020DecimalPowerDoesNotExceedRange(t *testing.T) {
	for _, r := range []float64{0.07, 0.7, 7, 70} {
		p := maxDecimalPower(r)
		if p > r {
			t.Fatalf("range=%v power=%v", r, p)
		}
	}
	if p := maxDecimalPower(100); p != 100 {
		t.Fatalf("exact power=%v", p)
	}
}
