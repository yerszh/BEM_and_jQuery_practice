package piscine

func splitSpace(str, separator string) []string {
	var words []string
	start := 0
	for i := 0; i < len(str); i++ {
		if str[i:i+len(separator)] == separator {
			words = append(words, str[start:i])
			start = i + len(separator)
			i += len(separator) - 1
		}
	}
	words = append(words, str[start:])
	return words
}

func ShoppingSummaryCounter(str string) map[string]int {
	counts := make(map[string]int)
	wordList := splitSpace(str, " ")

	for _, word := range wordList {
		counts[word]++
	}
	return counts
}
