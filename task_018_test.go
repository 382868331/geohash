package geohash
import "testing"
func TestTask018IntegerNeighborKeepsPrecision(t *testing.T){
	for _,bits:=range[]uint{15,25,40}{h:=EncodeIntWithPrecision(10,20,bits);all:=NeighborsIntWithPrecision(h,bits);for _,d:=range[]Direction{North,SouthWest}{got:=NeighborIntWithPrecision(h,bits,d);if got!=all[d]{t.Fatalf("bits=%d direction=%d got=%x want=%x",bits,d,got,all[d])}}}
}
