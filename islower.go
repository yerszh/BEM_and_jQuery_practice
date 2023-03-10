package piscine

func IsLower(s string) bool {
	flag := true
	for _, l := range s {
		if 'a' <= l && l <= 'z' {
			flag = true
		} else {
			flag = false
			break
		}
	}
	return flag
}
