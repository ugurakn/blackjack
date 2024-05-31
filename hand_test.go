package main

import (
	"fmt"
	"testing"

	"github.com/ugurakn/deck"
)

func Test_newHand(t *testing.T) {
	ply := player1
	h := newHand(ply)

	if h == nil {
		t.Fatalf("*hand is nil")
	}
	if h.owner != ply {
		t.Errorf("expected h.owner to be %v, got %v", ply, h.owner)
	}
	if h.cards == nil {
		t.Errorf("h.cards is nil. It must be initialized.")
	}
}

func Test_calc(t *testing.T) {
	h1 := newHand(player1)
	h1.cards = []deck.Card{
		{Suit: deck.Spade, Rank: deck.Two},
		{Suit: deck.Spade, Rank: deck.Five},
		{Suit: deck.Spade, Rank: deck.Eight},
	}
	h2 := newHand(player1)
	h2.cards = []deck.Card{
		{Suit: deck.Diamond, Rank: deck.Ten},
		{Suit: deck.Club, Rank: deck.Three},
		{Suit: deck.Heart, Rank: deck.J},
	}
	h3 := newHand(player1)
	h3.cards = []deck.Card{
		{Suit: deck.Spade, Rank: deck.J},
		{Suit: deck.Club, Rank: deck.Q},
		{Suit: deck.Diamond, Rank: deck.K},
	}
	h4 := newHand(player1)
	h4.cards = []deck.Card{
		{Suit: deck.Spade, Rank: deck.Ace},
		{Suit: deck.Club, Rank: deck.Q},
		{Suit: deck.Heart, Rank: deck.Ten},
	}
	h5 := newHand(player1)
	h5.cards = []deck.Card{
		{Suit: deck.Spade, Rank: deck.Three},
		{Suit: deck.Diamond, Rank: deck.Q},
		{Suit: deck.Heart, Rank: deck.Ace},
	}
	h6 := newHand(player1)
	h6.cards = []deck.Card{
		{Suit: deck.Heart, Rank: deck.Ace},
		{Suit: deck.Spade, Rank: deck.Three},
		{Suit: deck.Diamond, Rank: deck.Five},
		{Suit: deck.Heart, Rank: deck.Ace},
	}
	h7 := newHand(player1)
	h7.cards = []deck.Card{
		{Suit: deck.Club, Rank: deck.Ace},
		{Suit: deck.Spade, Rank: deck.Ace},
		{Suit: deck.Diamond, Rank: deck.Eight},
		{Suit: deck.Diamond, Rank: deck.Two},
	}
	h8 := newHand(player1)
	h8.cards = []deck.Card{
		{Suit: deck.Club, Rank: deck.Ace},
		{Suit: deck.Spade, Rank: deck.Ace},
		{Suit: deck.Diamond, Rank: deck.Ace},
		{Suit: deck.Heart, Rank: deck.Ace},
	}

	testCases := []struct {
		h      *hand
		expect int
	}{
		{h: h1, expect: 15},
		{h: h2, expect: 23},
		{h: h3, expect: 30},
		{h: h4, expect: 21},
		{h: h5, expect: 14},
		{h: h6, expect: 20},
		{h: h7, expect: 12},
		{h: h8, expect: 14},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("h%v", i+1), func(t *testing.T) {
			if val := tc.h.calc(); val != tc.expect {
				t.Errorf("expected hand val to be %v, got %v", tc.expect, val)
			}
		})
	}
}
