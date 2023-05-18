package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	for i := range deck {
		if i == 3 || i == 6 || i == 9 {
			player := i / 3
			fmt.Printf("Player %v: %v, %v, %v\n", player, deck[i-3], deck[i-2], deck[i-1])

		} else if i == 11 {
			fmt.Printf("Player %v: %v, %v, %v\n", 4, deck[i-2], deck[i-1], deck[i])
		}
	}
}
