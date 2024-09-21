package iteration

func Repeat(str string, count int) (result string) {
	result = ""
	for range count {
		result += str
	}

	return result
}
