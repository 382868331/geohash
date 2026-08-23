package geohash

import "testing"

func TestTask015SpreadSquashRoundTrip(t *testing.T) {
	for _, x := range []uint32{1, 3, 0x80000001, ^uint32(0)} {
		if got := squash(spread(x)); got != x {
			t.Fatalf("x=%x got=%x", x, got)
		}
	}
	if got := squash(0xaaaaaaaaaaaaaaaa); got != 0 {
		t.Fatalf("odd bits leaked: %x", got)
	}
}
