package geohash
import "testing"
func TestTask002BoxContainsEdges(t *testing.T){
	b:=Box{MinLat:-2,MaxLat:4,MinLng:10,MaxLng:20}
	for _,p:=range[][2]float64{{4,15},{1,20},{-2,10}}{if !b.Contains(p[0],p[1]){t.Fatalf("edge %#v rejected",p)}}
	if b.Contains(4.1,15){t.Fatal("outside latitude accepted")}
}
