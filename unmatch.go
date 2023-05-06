package piscine

func Unmatch(a []int) int {
	counter := 0

	for _, el := range a {
		counter = 0
		for _, n := range a {
			if n == el {
				counter++
			}
		}
		if counter%2 != 0 {
			return el
		}
	}
	return -1
}
