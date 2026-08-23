package geohash
import "testing"
func TestTask012EncodeRangeMidpoint(t *testing.T){
	if got:=encodeRange(0,90);got!=1<<31{t.Fatalf("latitude midpoint=%x",got)}
	if got:=encodeRange(0,180);got!=1<<31{t.Fatalf("longitude midpoint=%x",got)}
	if got:=encodeRange(-90,90);got!=0{t.Fatalf("minimum=%x",got)}
}
