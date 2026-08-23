package geohash
import "testing"
func TestTask003BoxCenterAxisOrder(t *testing.T){
	b:=Box{MinLat:-10,MaxLat:20,MinLng:100,MaxLng:140};lat,lng:=b.Center()
	if lat!=5||lng!=120{t.Fatalf("center=(%v,%v)",lat,lng)}
	b=Box{MinLat:0,MaxLat:2,MinLng:-8,MaxLng:-2};lat,lng=b.Center()
	if lat!=1||lng!=-5{t.Fatalf("second center=(%v,%v)",lat,lng)}
}
