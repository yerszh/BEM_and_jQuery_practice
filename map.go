package piscine

func Map(f func(int) bool, a []int) []bool {
	result := []bool{}

	for _, el := range a {
		result = append(result, f(el))
	}
	return result
}
