package main

import (
	"testing"

	"github.com/ugurakn/deck"
)

func Test_deal(t *testing.T) {
	d := deck.New()
	dOrigLen := len(d)
	h := &hand{owner: player1, cards: make([]deck.Card, 0)}

	d = deal(d, h)

	if len(d) != dOrigLen-1 {
		t.Errorf("expected deck len %v, got %v", dOrigLen-1, len(d))
	}
	if d[0].Suit != deck.Spade || d[0].Rank != deck.Two {
		t.Errorf("expected d[0] to be an Two of Spades, found %v", d[0])
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
	d := deck.New()
	dOrigLen := len(d)

	dh := newHand(dealer)
	players := []*hand{newHand(player1), newHand(player2)}

	for i := 0; i < 2; i++ {
		for _, p := range players {
			d = deal(d, p)
		}
		d = deal(d, dh)
	}

	if expectLen := dOrigLen - 6; len(d) != expectLen {
		t.Errorf("expected deck length to be %v after dealing, got %v", expectLen, len(d))
	}
	for _, p := range append(players, dh) {
		if len(p.cards) != 2 {
			t.Errorf("expected the len of %v's cards to be 2, found %v", p.owner, len(p.cards))
		}
	}
}

// func Test_player_string(t *testing.T) {
// 	fmt.Println(dealer)
// 	fmt.Println(player1)
// 	fmt.Println(player2)
// }
