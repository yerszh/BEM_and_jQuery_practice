package piscine

func ToLower(s string) string {
	result := ""
	for _, l := range s {
		if 'A' <= l && l <= 'Z' {
			result = result + string(l+32)
		} else {
			result = result + string(l)
		}
	}
	return result
}
