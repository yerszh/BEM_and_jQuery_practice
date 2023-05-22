package piscine

func StringToIntSlice(str string) []int {
	if str == "" {
		return []int(nil)
	}
	result := []int{}

	for _, r := range str {
		result = append(result, int(r))
	}
	return result
}
