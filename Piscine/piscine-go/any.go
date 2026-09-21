package piscine

func Any(f func(string) bool, a []string) bool {
	for _, v := range a { // Iterate through the slice 'a'
		if f(v) { // Apply 'f' on each string 'v'
			return true // If any string satisfies the condition, return true
		}
	}
	return false // If no string satisfies the condition, return false
}
