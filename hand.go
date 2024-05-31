package main

import "github.com/ugurakn/deck"

type hand struct {
	owner player
	cards []deck.Card
	val   int
}

// newHand creates and returns
// a new *hand with owner o
func newHand(o player) *hand {
	return &hand{owner: o, cards: make([]deck.Card, 0), val: 0}
}

// calc calculates the current value of cards in a hand.
func (h *hand) calc() int {
	var total int
	var aces int
	for _, c := range h.cards {
		if c.Rank == deck.Ten || c.Rank == deck.J || c.Rank == deck.Q || c.Rank == deck.K {
			total += 10
		} else if c.Rank == deck.Ace {
			aces++
		} else { // c.Rank in [2,9]
			total += int(c.Rank)
		}
	}
	// add 1 per ace to total
	total += aces
	// turn one ace to 11 if total wouldn't bust
	if total <= 11 {
		total += 10
	}
	h.val = total
	return total
}
