package arrays

func Sum(numbers [5]int) int {
	var sum int
	for _, val := range numbers {
		sum += val
	}
	return sum
}
