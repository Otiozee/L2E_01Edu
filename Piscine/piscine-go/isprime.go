package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false // 1 and negative numbers are not prime
	}
	if nb == 2 {
		return true // 2 is the only even prime number
	}
	if nb%2 == 0 {
		return false // Return false for even numbers greater than 2
	}

	// Check divisibility from 3 to the square root of nb
	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
