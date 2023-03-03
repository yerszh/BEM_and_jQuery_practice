package piscine

func AlphaCount(s string) int {
	counter := 0
	for _, l := range []rune(s) {
		if 'a' <= l && l <= 'z' || 'A' <= l && l <= 'Z' {
			counter++
		}
	}
	return counter
}
