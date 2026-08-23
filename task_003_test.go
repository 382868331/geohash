package geohash
import "testing"
func TestTask003BoxCenterAxisOrder(t *testing.T){
	b:=Box{MinLat:-10,MaxLat:20,MinLng:100,MaxLng:140};lat,lng:=b.Center()
	if lat!=5||lng!=120{t.Fatalf("center=(%v,%v)",lat,lng)}
}
