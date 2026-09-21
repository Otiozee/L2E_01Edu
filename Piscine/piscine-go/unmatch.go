package piscine

func Unmatch(a []int) int {
	n := len(a)
	for i := 0; i < n; i++ { // Outer loop: this iterate through each element in the slice
		count := 0
		for j := 0; j < n; j++ { // Inner loop: this count how many times current element appears
			if a[i] == a[j] {
				count++ // If elements match, increment counter
			}
		}
		if count%2 != 0 {
			return a[i]
		}
	}
	return -1 // If all elements have even counts, return -1
}
