package main

import (
	"fmt"

	"github.com/ugurakn/deck"
)

type winState int

// possible cases: undecided, lost, lost (bust), push (i.e. draw), won, won (blackjack)
const (
	undecided winState = iota
	lost
	bust
	push
	win
	winbj
)

type hand struct {
	owner player
	cards []deck.Card
	// if the hand is a natural blackjack
	bjack bool
	// if the player is human or AI
	human bool
	winState
}

func (h *hand) String() string {
	s := fmt.Sprintf("(%v) hand: ", h.owner)

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
	if total <= 11 && aces > 0 {
		total += 10
	}
	return total
}

// set winState for h by comparison to
// dealer's hand dh
func (h *hand) setWinState(dh *hand) {
	if dh.bjack {
		if h.bjack {
			h.winState = push
		} else {
			h.winState = lost
		}
		return
	}

	if h.bjack {
		h.winState = winbj
		return
	}

	val := h.value()
	dhVal := dh.value()

	if val > 21 {
		h.winState = bust
		return
	}
	if val < dhVal {
		h.winState = lost
	} else if val > dhVal {
		h.winState = win
	} else {
		h.winState = push
	}
}

// newHand creates and returns
// a new *hand with owner o
func newHand(o player, human bool) *hand {
	return &hand{owner: o, cards: make([]deck.Card, 0), human: human}
}
