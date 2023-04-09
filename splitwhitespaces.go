package piscine

func SplitWhiteSpaces(s string) []string {
	result := []string{}
	word := ""

	for _, r := range s {
		if r == ' ' || r == '\n' || r == '\t' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(r)
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
