package piscine

// isPrime checks if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// FindNextPrime returns the first prime number >= nb
func FindNextPrime(nb int) int {
	for !isPrime(nb) {
		nb++
	}
	return nb
}
