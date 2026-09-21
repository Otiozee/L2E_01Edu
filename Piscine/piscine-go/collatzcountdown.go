package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	steps := 0
	current := start
	for current != 1 {
		if current%2 == 0 {
			current /= 2
		} else {
			current = current*3 + 1
		}
		steps++
	}
	return steps
}
