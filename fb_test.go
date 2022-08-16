package fizzbuzz
import "testing"

var tests = map[string]int{
	"FIZZ":     3,
	"BUZZ":     5,
}

func TestFizzbuzz(t *testing.T){
	for key,res := range tests {
		if Fizzbuzz(res) !=key {
			t.Errorf("Failed on %s != %v", key, res)
		}
	}
}
