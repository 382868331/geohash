package geohash
import "testing"
func TestTask016RoundedPointStaysInside(t *testing.T){
	for _,b:=range[]Box{{1.23,1.29,2.34,2.39},{-1.29,-1.23,-2.39,-2.34}}{lat,lng:=b.Round();if !b.Contains(lat,lng){t.Fatalf("box=%#v rounded=(%v,%v)",b,lat,lng)}}
}
