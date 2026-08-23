package geohash
import "testing"
func TestTask010EncodingPrecisionLength(t *testing.T){
	for _,n:=range[]uint{1,5,11}{h:=EncodeWithPrecision(57.64911,10.40744,n);if len(h)!=int(n){t.Fatalf("precision=%d hash=%q",n,h)}}
}
