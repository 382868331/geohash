package geohash
import "testing"
func TestTask006Base32EncodeBitStride(t *testing.T){
	for _,x:=range[]uint64{0,31,32,1024,0x12345}{s:=base32encoding.Encode(x);if base32encoding.Decode(s)!=x{t.Fatalf("x=%x encoded=%q decoded=%x",x,s,base32encoding.Decode(s))}}
}
