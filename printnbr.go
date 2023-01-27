package piscine

import "github.com/01-edu/z01"

func intToRune(num int) rune {
	if num < 0 {
		num *= -1
	}
	switch num {
	case 0:
		return '0'
	case 1:
		return '1'
	case 2:
		return '2'
	case 3:
		return '3'
	case 4:
		return '4'
	case 5:
		return '5'
	case 6:
		return '6'
	case 7:
		return '7'
	case 8:
		return '8'
	default:
		return '9'
	}
}

func PrintNbr(n int) {
	neg := 1

	if n < 0 {
		neg = -1
		z01.PrintRune('-')
	}
	if n >= 10 || n <= -10 {
		PrintNbr(n / 10 * neg)
	}
	z01.PrintRune(intToRune(n % 10))
}
