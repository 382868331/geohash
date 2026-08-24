package geohash

import "testing"

func TestTask019StringConversionPreservesFiveBitDigits(t *testing.T) {
	value, bits := ConvertStringToInt("ezs42")
	if bits != 25 {
		t.Fatalf("precision = %d bits, want 25", bits)
	}
	if value != 0xdfe082 {
		t.Fatalf("decoded value = %#x, want %#x", value, uint64(0xdfe082))
	}
}
