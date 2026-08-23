package geohash
import "testing"
func TestTask008PrecisionAxisRanges(t *testing.T){
	lat,lng:=errorWithPrecision(2);if lat!=90||lng!=180{t.Fatalf("errors=(%v,%v)",lat,lng)}
	lat,lng=errorWithPrecision(4);if lat!=45||lng!=90{t.Fatalf("four-bit errors=(%v,%v)",lat,lng)}
}
