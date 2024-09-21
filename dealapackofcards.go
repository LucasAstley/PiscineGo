package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	for i := 1; i <= 4; i++ {
		fmt.Printf("Player %d", i)
		fmt.Printf(": ")
		for j := 0; j < 3; j += 1 {
			fmt.Printf("%d", deck[0])
			if j < 2 {
				fmt.Printf(", ")
			}
			deck = deck[1:]
		}
		z01.PrintRune('\n')
	}
}
