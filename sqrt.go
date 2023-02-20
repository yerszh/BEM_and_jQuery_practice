package piscine

func Sqrt(nb int) int {
	if nb > 1 {
		for i := 0; i < nb; i++ {
			if i*i == nb {
				return i
			}
		}
	} else if nb == 1 {
		return 1
	}
	return 0
}
