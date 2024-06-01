package main

import (
	"fmt"

	"github.com/ugurakn/deck"
)

type hand struct {
	owner player
	cards []deck.Card
	bust  bool
}

func (h *hand) String() string {
	// example:
	// Player's Hand: 'Ace of Spades', 'Five of Diamonds' -> value: 16
	s := fmt.Sprintf("%s's hand: ", h.owner)

	var hiddenSlice int
	if h.owner == dealer {
		s += "[face-down], "
		hiddenSlice = 1
	}

	cards := h.cards[hiddenSlice:]
	for i, c := range cards {
		s += fmt.Sprintf("'%s'", c)
		if i != len(cards)-1 {
			s += ", "
		}
	}
	if h.owner != dealer {
		s += fmt.Sprintf(" -> value: %v", h.value())
	}
	return s
}

// calc calculates the current value of cards in a hand.
func (h *hand) value() int {
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
	return total
}

// newHand creates and returns
// a new *hand with owner o
func newHand(o player) *hand {
	return &hand{owner: o, cards: make([]deck.Card, 0)}
}
