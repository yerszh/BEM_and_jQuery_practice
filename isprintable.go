package piscine

func IsPrintable(s string) bool {
	flag := true
	for _, l := range s {
		if l >= 0 && l <= 31 {
			return !flag
		}
	}
	return flag
}
