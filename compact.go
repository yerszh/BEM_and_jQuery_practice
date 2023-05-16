package piscine

func Compact(ptr *[]string) int {
	counter := 0

	arr := []string{}

	for _, el := range *ptr {
		if el != "" {
			arr = append(arr, el)
			counter++
		}
	}

	*ptr = arr

	return counter
}
