package piscine

func Capitalize(s string) string {
	runeArr := []rune(s)

	firstLetter := true

	for i, l := range runeArr {
		if IsAlpha(string(l)) && firstLetter {
			if IsLower(string(l)) {
				runeArr[i] -= 32
			}
			firstLetter = false

		} else if IsUpper(string(l)) {
			runeArr[i] += 32
		} else if !IsAlpha(string(runeArr[i])) {
			firstLetter = true
		}
	}
	return string(runeArr)
}
