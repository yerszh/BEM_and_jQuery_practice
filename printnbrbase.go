package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 || containsPlusMinus(base) || !hasUniqueCharacters(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	sign := ""
	if nbr < 0 {
		sign = "-"
		nbr = -nbr
	}

	result := ""
	for nbr > 0 {
		remainder := nbr % len(base)
		result = string(base[remainder]) + result
		nbr = nbr / len(base)
	}

	for _, i := range sign + result {
		z01.PrintRune(i)
	}
}

func containsPlusMinus(s string) bool {
	for _, char := range s {
		if char == '+' || char == '-' {
			return true
		}
	}
	return false
}

func hasUniqueCharacters(s string) bool {
	seen := make(map[rune]bool)
	for _, char := range s {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}
