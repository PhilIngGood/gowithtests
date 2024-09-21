package slice

func Sum(list []int) (result int) {
	for _, value := range list {
		result += value
	}
	return result
}
