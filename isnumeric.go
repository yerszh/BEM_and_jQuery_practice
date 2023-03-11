package piscine

func IsNumeric(s string) bool {
	flag := true
	for _, l := range s {
		if '0' <= l && l <= '9' {
			flag = true
		} else {
			flag = false
			break
		}
	}
	return flag
}
