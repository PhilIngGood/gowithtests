package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 7)
	expected := strings.Repeat("a", 7)
	if repeated != expected {
		t.Errorf("expected %q, got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	res := Repeat("q", 10)
	fmt.Println(res)
	//  Output: qqqqqqqqqq
}

func BenchmarkRepeat(b *testing.B) {
	for range b.N {
		Repeat("a", 5)
	}
}
