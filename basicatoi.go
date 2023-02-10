package piscine

func BasicAtoi(s string) int {
	bs := []byte(s)

	result := 0
	multi := 1
	for i := len(bs) - 1; i >= 0; i-- {

		result += ((int(bs[i]) - 48) * multi)
		multi *= 10
	}
	return result
}
