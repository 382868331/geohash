package geohash
import "testing"
func TestTask011NeighborDirectionMapping(t *testing.T){
	h:="u4pruy";all:=Neighbors(h)
	for _,d:=range[]Direction{North,East,South,West}{if got:=Neighbor(h,d);got!=all[d]{t.Fatalf("direction=%d got=%q want=%q",d,got,all[d])}}
}
