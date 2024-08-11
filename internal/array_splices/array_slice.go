package array_slice

func Sum(numbers []int) int {
	result := 0
	for _, v := range numbers {
		result += v
	}
	return result
}
