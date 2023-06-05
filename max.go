package piscine

func Max(a []int) int {
	result := a[0]

	for _, el := range a {
		if el > result {
			result = el
		}
	}
	return result
}
