package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) > 2 {
		first := args[0]
		op := args[1][0]
		second := args[2]
		if !(IsNumeric(first)) || !(IsNumeric(second)) {
			return
		}
		a := parse_int(first)
		b := parse_int(second)
		if op == '%' && b == 0 {
			printStr("No modulo by 0")
			z01.PrintRune('\n')

			return
		} else if op == '/' && b == 0 {
			printStr("No division by 0")
			z01.PrintRune('\n')
			return
		}
		result := 0
		if a > 9223372036854775750 || b > 9223372036854775750 {
			return
		}
		switch op {
		case '*':
			result = a * b
			if a == 0 || result/a == b {
				PrintNbr(result)
				z01.PrintRune('\n')
			}
		case '/':
			result = a / b
			PrintNbr(result)
			z01.PrintRune('\n')
		case '%':
			result = a % b
			PrintNbr(result)
			z01.PrintRune('\n')
		case '-':
			result = a - b
			if (result < a) && (b > 0) {
				PrintNbr(result)
				z01.PrintRune('\n')
			}
		case '+':
			result = a + b
			if (result > a) && (b > 0) {
				PrintNbr(result)
				z01.PrintRune('\n')
			}
		}
	}
}

func IsNumeric(s string) bool {
	runes := []rune(s)
	for i := range runes {
		if i == 0 && runes[0] == '-' {
			continue
		}
		if !(runes[i] >= 48 && runes[i] <= 57) {
			return false
		}
	}
	return true
}

func parse_int(s string) int {
	minus := false
	result := 0
	for i, c := range s {
		if s[0] == '-' && i == 0 {
			minus = true
			continue
		}
		result *= 10
		result += int(c - '0')
	}
	if minus {
		result = -result
	}
	return result
}

func printStr(str string) {
	for _, i := range str {
		z01.PrintRune(i)
	}
}

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
