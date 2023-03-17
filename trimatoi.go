package piscine

func TrimAtoi(s string) int {
	isPlus := true
	hadNumber := false
	var intArr []int
	for _, l := range s {
		if '0' <= l && l <= '9' {
			if l == '0' && !hadNumber {
				continue
			}

			intArr = append(intArr, int(l))
			hadNumber = true
		} else if l == '-' && !hadNumber {
			isPlus = false
		}
	}
	var outInt int
	op := 1
	for i := len(intArr) - 1; i >= 0; i-- {
		outInt += int(intArr[i]-'0') * op
		op *= 10
	}
	if !isPlus {
		outInt *= -1
	}

	return outInt
}
