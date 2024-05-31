package arraysslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers { // _ = idx, number = value
		sum += number
	}
	return sum
}
