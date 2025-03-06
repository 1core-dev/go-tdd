package arrays

func Sum(numbers []int) int {
	var sum int
	for _, val := range numbers {
		sum += val
	}
	return sum
}
