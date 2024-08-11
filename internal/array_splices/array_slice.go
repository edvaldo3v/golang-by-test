package array_slice

func Sum(numbers []int) int {
	result := 0
	for _, v := range numbers {
		result += v
	}
	return result
}

func SumAll(numbersSlices ...[]int) []int {
	var sums []int

	for _, numbers := range numbersSlices {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllRest(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			end := numbers[1:]
			sums = append(sums, Sum(end))
		}
	}

	return sums
}
