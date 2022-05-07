package arrays

import "testing"

func TestSum(t *testing.T) {

	t.Run("sum is 5 number, sum=15", func(t *testing.T) {

		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("sum is 5 number, sum=25", func(t *testing.T) {

		numbers := [5]int{1, 2, 3, 4, 15}

		got := Sum(numbers)
		want := 25

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
