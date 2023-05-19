package piscine

func JumpOver(str string) string {
	result := ""

	for i := range str {
		if i%3 == 2 {
			result += string(str[i])
		}
	}
	result += string("\n")
	return result
}
