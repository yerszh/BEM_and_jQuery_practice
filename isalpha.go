package piscine

func IsAlpha(s string) bool {
	flag := true
	for _, l := range s {
		if '0' <= l && l <= '9' || 'a' <= l && l <= 'z' || 'A' <= l && l <= 'Z' {
			flag = true
		} else {
			flag = false
			break
		}
	}
	return flag
}
