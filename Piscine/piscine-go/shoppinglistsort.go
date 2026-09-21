package piscine

func ShoppingListSort(slice []string) []string {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if len(slice[j]) < len(slice[minIndex]) {
				minIndex = j
			}
		}
		slice[i], slice[minIndex] = slice[minIndex], slice[i]
	}
	return slice
}
