package arrays

func Sum(numbers []int) int {
	var sum int
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	numbersLen := len(numbersToSum)
	sums := make([]int, numbersLen)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
