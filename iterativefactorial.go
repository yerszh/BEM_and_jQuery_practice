package piscine

func IterativeFactorial(nb int) int {
	if nb >= 0 && nb <= 20 {
		factorial := 1
		for i := 1; i <= nb; i++ {
			factorial *= i
		}
		return factorial
	} else {
		return 0
	}
}
