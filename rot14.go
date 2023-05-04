package piscine

func Rot14(s string) string {
	result := ""

	for _, r := range s {
		if 'A' <= r && r <= 'Z' {
			result += string('A' + (r-'A'+14)%26)
		} else if 'a' <= r && r <= 'z' {
			result += string('a' + (r-'a'+14)%26)
		} else {
			result += string(r)
		}
	}

	return result
}
