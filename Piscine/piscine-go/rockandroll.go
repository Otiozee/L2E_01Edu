package piscine

func RockAndRoll(n int) string {
	// If the number is negative
	if n < 0 {
		return "error: number is negative\n"
	}

	// If divisible by both 2 and 3
	if n%2 == 0 && n%3 == 0 {
		return "rock and roll\n"
	}

	// If divisible by 2
	if n%2 == 0 {
		return "rock\n"
	}

	// If divisible by 3
	if n%3 == 0 {
		return "roll\n"
	}

	// If not divisible by 2 or 3
	return "error: non divisible\n"
}
