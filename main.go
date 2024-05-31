package main

import (
	"fmt"

	"github.com/ugurakn/deck"
)

const initialDealSize = 2

func main() {
	d := deck.New(deck.Shuffle)

	// this is for only a single player
	p1Hand := newHand(player1)
	dHand := newHand(dealer)

	players := []*hand{p1Hand}

	// game loop

	var done bool
	for !done {
		// initial deal phase TODO
		for i := 0; i < initialDealSize; i++ {
			for _, p := range players {
				d = deal(d, p)
			}
			d = deal(d, dHand)
		}

		fmt.Println("dealer's hand:", dHand.cards)
		fmt.Println("player's hand:", players[0].cards)
		fmt.Println(len(d), d[0])

		// finish game
		done = true
	}
}
