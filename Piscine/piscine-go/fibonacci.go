package piscine

func Fibonacci(index int) int {
	// Base case for invalid negative indices
	if index < 0 {
		return -1
	}
	// Base case for Fibonacci(0)
	if index == 0 {
		return 0
	}
	// Base case for Fibonacci(1)
	if index == 1 {
		return 1
	}
	// Recursive call for Fibonacci(index-1) + Fibonacci(index-2)
	return Fibonacci(index-1) + Fibonacci(index-2)
}
