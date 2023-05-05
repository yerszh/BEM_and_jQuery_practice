package piscine

import "github.com/01-edu/z01"

func printInt(x int) {
	secondDigit := (x % 10)
	firstDigit := (x / 10)

	z01.PrintRune(rune(firstDigit) + 48)
	z01.PrintRune(rune(secondDigit) + 48)
}

func DescendComb() {
	for i := 99; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			printInt(i)
			z01.PrintRune(' ')
			printInt(j)
			if i == 1 && j == 0 {
				break
			}
			z01.PrintRune(',')
			z01.PrintRune(' ')

		}
	}
}
