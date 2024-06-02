package main

import "github.com/ugurakn/deck"

// type rValue struct {
// 	r   deck.Rank
// 	val int
// }

// var (
// 	ace   = rValue{deck.Ace, 11}
// 	two   = rValue{deck.Two, 2}
// 	three = rValue{deck.Three, 3}
// 	four  = rValue{deck.Four, 4}
// 	five  = rValue{deck.Five, 5}
// 	six   = rValue{deck.Six, 6}
// 	seven = rValue{deck.Seven, 7}
// 	eigth = rValue{deck.Eight, 8}
// 	nine  = rValue{deck.Nine, 9}
// 	ten   = rValue{deck.Ten, 10}
// 	faceJ = rValue{deck.J, 10}
// 	faceQ = rValue{deck.Q, 10}
// 	faceK = rValue{deck.K, 10}
// )

type player int

const (
	dealer player = iota
	player1
	player2
)

//go:generate stringer -type=player

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

// checkBJ is called for each player's and dealer's
// hand to see whether it is a natural blackjack.
// If it is, sets h.bjack to true
// and returns true.
func checkBJ(h *hand) bool {
	if h.value() == 21 {
		h.bjack = true
		return true
	}
	return false
}
