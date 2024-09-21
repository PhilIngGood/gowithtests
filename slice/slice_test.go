package slice

import "testing"

// Test sum on Array
func TestSum(t *testing.T) {
	t.Run("collection of 	 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		got := Sum(numbers)
		want := len(numbers) * (len(numbers) + 1) / 2

		if want != got {
			t.Errorf("want %d, got %d, given %v", want, got, numbers)
		}
	})
}
