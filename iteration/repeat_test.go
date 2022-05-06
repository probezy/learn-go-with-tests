package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 6)
	want := "aaaaaa"

	if got != want {
		t.Errorf("want '%q' but got '%q'", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Repeat("a", 6)
	}
}

func ExampleRepeat() {
	output := Repeat("a", 5)
	fmt.Println(output)
	//Output: aaaaa
}
