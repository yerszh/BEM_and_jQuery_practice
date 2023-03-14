package piscine

func ToUpper(s string) string {
	result := ""
	for _, l := range s {
		if 'a' <= l && l <= 'z' {
			result = result + string(l-32)
		} else {
			result = result + string(l)
		}
	}
	return result
}
