package main

import (
	"testing"

	"github.com/ugurakn/deck"
)

func Test_deal(t *testing.T) {
	sh := new(shoe)
	sh.cards = deck.New()
	sh.initSize = len(sh.cards)
	h := &hand{owner: player1, cards: make([]deck.Card, 0)}

	deal(sh, h)

	if len(sh.cards) != sh.initSize-1 {
		t.Errorf("expected deck len %v, got %v", sh.initSize-1, len(sh.cards))
	}
	if sh.cards[0].Suit != deck.Spade || sh.cards[0].Rank != deck.Two {
		t.Errorf("expected d[0] to be an Two of Spades, found %v", sh.cards[0])
	}
	if len(h.cards) != 1 {
		t.Errorf("expected len(h.cards) to be 1, found %v", len(h.cards))
	}
	if h.cards[0].Suit != deck.Spade || h.cards[0].Rank != deck.Ace {
		t.Errorf("Expected the card at h.cards[0] to be an Ace of Spades, found %v", h.cards[0])
	}
}

// tests the initial deal 2 to players & dealer.
func Test_initial_deal(t *testing.T) {
	sh := new(shoe)
	sh.cards = deck.New()
	sh.initSize = len(sh.cards)

	dh := newHand(dealer)
	players := []*hand{newHand(player1), newHand(player2)}

	for i := 0; i < 2; i++ {
		for _, p := range players {
			deal(sh, p)
		}
		deal(sh, dh)
	}

	if expectLen := sh.initSize - 6; len(sh.cards) != expectLen {
		t.Errorf("expected deck length to be %v after dealing, got %v", expectLen, len(sh.cards))
	}
	for _, p := range append(players, dh) {
		if len(p.cards) != 2 {
			t.Errorf("expected the len of %v's cards to be 2, found %v", p.owner, len(p.cards))
		}
	}
}
