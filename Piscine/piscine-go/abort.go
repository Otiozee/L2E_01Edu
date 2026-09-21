package piscine

func Abort(a, b, c, d, e int) int {
	// Create a slice to hold the numbers
	numbers := []int{a, b, c, d, e}

	// Sort the slice in ascending order
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	// The median is the third element in the sorted slice
	return numbers[2]
}
