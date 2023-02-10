package piscine

func BasicAtoi2(s string) int {
	bs := []byte(s)

	result := 0
	multi := 1
	tmp := 0
	for i := len(bs) - 1; i >= 0; i-- {
		tmp = (int(bs[i]) - 48)
		if tmp < 0 || tmp > 9 {
			return 0
		}
		result += (tmp * multi)
		multi *= 10
	}
	return result
}
