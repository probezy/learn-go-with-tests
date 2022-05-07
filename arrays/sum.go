package arrays

// Sum
func Sum(numbers []int) (sum int) {

	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sum := Sum(numbers)
		sums = append(sums, sum)
	}
	return
}
