package slice

import (
	"slices"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	t.Run("Testing SumAll", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		//if !reflect.DeepEqual(got, want) {
		//	t.Errorf("got %v, want %v", got, want)
		//}
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("Testing TestSumAllTails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9})
		want := []int{5, 9}
		checkSums(t, got, want)
	})
	t.Run("Testing sum empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
