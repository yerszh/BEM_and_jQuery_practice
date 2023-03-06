package piscine

func IsUpper(s string) bool {
	flag := true
	for _, l := range s {
		if 'A' <= l && l <= 'Z' {
			flag = true
		} else {
			flag = false
			break
		}
	}
	return flag
}
