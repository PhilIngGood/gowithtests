package slice

func Sum(list []int) (result int) {
	for _, value := range list {
		result += value
	}
	return result
}

func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))
	for i, slice := range numbersToSum {
		sums[i] = Sum(slice)
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))
	for i, slice := range numbersToSum {
		if len(slice) == 0 {
			sums[i] = 0
		} else {
			sums[i] = Sum(slice[1:])
		}
	}
	return sums
}
