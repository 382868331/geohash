package geohash
import "testing"
func TestTask004ValidateMaximumLength(t *testing.T){
	if err:=Validate("0123456789bc");err!=nil{t.Fatalf("12 chars rejected: %v",err)}
	if err:=Validate("0123456789bcd");err==nil{t.Fatal("13 chars accepted")}
}
