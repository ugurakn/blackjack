package main

import "github.com/ugurakn/deck"

type player int

const (
	dealer player = iota
	player1
	player2
)

//go:generate stringer -type=player

type hand struct {
	owner player
	cards []deck.Card
}

// newHand creates and returns a new *hand
// with owner o and empty cards
func newHand(o player) *hand {
	return &hand{owner: o, cards: make([]deck.Card, 0)}
}

// deal pops one card off the top of d
// and appends it to h.cards.
// Returns modified d.
// Panics if len(d) == 0
func deal(d []deck.Card, h *hand) []deck.Card {
	if len(d) == 0 {
		panic("can't deal from empty deck")
	}

	h.cards = append(h.cards, d[0])
	d = d[1:]
	return d
}

// Draw draws the top card from d
// and returns the modified d and the card.
// func Draw(d []deck.Card) ([]deck.Card, deck.Card) {
// 	if len(d) == 0 {
// 		panic("Draw: can't draw from empty deck.")
// 	}
// 	c := d[0]
// 	d = d[1:]
// 	return d, c
// }
