package piscine

func Map(f func(int) bool, a []int) []bool {
	var result []bool // Create an empty slice to store the results.

	for _, v := range a { // Loop through each element of the slice 'a'
		result = append(result, f(v)) // Apply 'f' on 'v' and append the result to 'result'
	}

	return result // Return the final slice
}
