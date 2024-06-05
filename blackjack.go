package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/ugurakn/deck"
)

// shoe represents the game deck
// from which cards are dealt
type shoe struct {
	cards    []deck.Card
	initSize int
}

// type player int

// const (
// 	dealer player = iota
// 	player1
// 	player2
// )

// deal pops one card off the top of d
// and appends it to h.cards.
// Returns modified d.
// Panics if len(d) == 0
func deal(sh *shoe, h *hand) {
	if len(sh.cards) == 0 {
		panic("can't deal from empty deck")
	}

	h.cards = append(h.cards, sh.cards[0])
	sh.cards = sh.cards[1:]
}

// hit is a player action that adds a new card
// to a hand (using deal).
// returns if the player busted.
func hit(sh *shoe, h *hand) bool {
	deal(sh, h)
	return h.value() > 21
}

// if hand lost: owner loses the bet amount.
// if push: nothing happens.
// if won: the owner gains 1:1.
// if won with bjack: owner gains 3:2.
// NB bet amounts are not pulled out of player
// purses on betting. A bet of 100 will net
// a player +100 in their purse if they win.
// returns the amount won/lost.
func payout(h *hand) int {
	ws := h.winState
	if ws == lost || ws == bust {
		h.owner.purse -= h.bet
		if h.owner.purse < 0 {
			panic("payout: player purse can't go below 0.")
		}
		return h.bet * -1
	} else if ws == win {
		h.owner.purse += h.bet
		return h.bet
	} else if ws == winbj {
		amount := int(float32(h.bet) * 1.5)
		h.owner.purse += amount
		return amount
	}
	// push
	return 0
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

// playTurn will let a player play his turn.
func playTurn(sh *shoe, p *hand) {
	firstTurn := true
	var done bool
	for !done {
		fmt.Println(p)

		// get player input
		if firstTurn {
			fmt.Printf("(%v) (h)it, (d)ouble down or (s)tand: ", p.owner)
		} else {
			fmt.Printf("(%v) (h)it, (s)tand: ", p.owner)
		}
		var in string
		if p.owner.isHuman {
			in = getInput()
		} else {
			in = getAIInput()
		}
		switch strings.ToLower(in) {
		case "h": // hit
			bust := hit(sh, p)
			val := p.value()
			fmt.Printf("(%v) HIT: '%v' new value:%v\n", p.owner, p.cards[len(p.cards)-1], val)
			if val == 21 { // auto-stand on 21
				done = true
			}
			if bust {
				done = true
				fmt.Printf("(%v) BUST!\n", p.owner)
			}

		case "d": // double-down
			if !firstTurn {
				fmt.Println("can't double-down now.")
				continue
			}
			// double bet
			p.bet *= 2
			fmt.Printf("(%v) bet doubled (new bet: %v)\n", p.owner, p.bet)
			// hit once, then stand
			bust := hit(sh, p)
			fmt.Printf("(%v) HIT: %v. new value:%v\n", p.owner, p.cards[len(p.cards)-1], p.value())
			if bust {
				fmt.Printf("(%v) BUST!\n", p.owner)
			}
			done = true

		case "s": // stand
			fmt.Printf("(%v) STAND\n", p.owner)
			done = true

		default:
			continue
		}

		if firstTurn {
			firstTurn = false
		}
		time.Sleep(time.Millisecond * 500)
	}
}

// dealer will hit on 16 or less,
// stand on > 16
func playDealer(sh *shoe, h *hand) {
	var done bool
	for !done {
		if h.value() <= 16 {
			bust := hit(sh, h)
			fmt.Printf("(%v) HIT: '%v' new value:%v\n", h.owner, h.cards[len(h.cards)-1], h.value())
			if bust {
				done = true
				fmt.Printf("(%v) BUST!\n", h.owner)
			}
		} else {
			fmt.Printf("(%v) STAND\n", h.owner)
			done = true
		}
		time.Sleep(time.Millisecond * 500)
	}
}

// Get input from human player.
func getInput() string {
	var in string
	fmt.Scanf("%s\n", &in)
	return in
}

// To be implemented
func getAIInput() string {
	// TODO
	return "AI"
}
