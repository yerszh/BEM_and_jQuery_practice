package piscine

func Split(s, sep string) []string {
	var result []string
	startIndex := 0
	separatorLen := len(sep)

	for i := 0; i < len(s); i++ {
		if i+separatorLen <= len(s) && s[i:i+separatorLen] == sep {
			result = append(result, s[startIndex:i])
			startIndex = i + separatorLen
			i += separatorLen - 1
		}
	}

	if startIndex < len(s) {
		result = append(result, s[startIndex:])
	}

	return result
}
