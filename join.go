package piscine

func Join(strs []string, sep string) string {
	result := ""

	for i, str := range strs {
		result += str

		if i == len(strs)-1 {
			break
		}
		result += sep
	}

	return result
}
