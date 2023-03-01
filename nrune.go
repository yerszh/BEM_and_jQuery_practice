package piscine

func NRune(s string, n int) rune {
	if len(s) >= n && n > 0 {
		return []rune(s)[n-1]
	} else {
		return 0
	}
}
