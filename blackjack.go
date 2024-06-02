package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ugurakn/deck"
)

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

// shoe represents the game deck
// from which cards are dealt
type shoe struct {
	cards    []deck.Card
	initSize int
}

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
	var done bool
	for !done {
		fmt.Println(p)

		// get player input
		fmt.Printf("(%v) (h)it or (s)tand: ", p.owner)
		var in string
		if p.human {
			in = getInput()
		} else {
			in = getAIInput()
		}
		switch strings.ToLower(in) {
		case "h": // hit
			bust := hit(sh, p)
			val := p.value()
			fmt.Printf("(%v) HIT: %v. new value:%v\n", p.owner, p.cards[len(p.cards)-1], val)
			if val == 21 { // auto-stand on 21
				done = true
			}
			if bust {
				done = true
				fmt.Printf("(%v) BUST!\n", p.owner)
			}

		case "s": // stand
			fmt.Printf("(%v) STAND\n", p.owner)
			done = true

		default:
			fmt.Println("unknown input:", in)
			os.Exit(1)
		}
		time.Sleep(time.Millisecond * 500)
	}
}

// Get input from human player.
// Valid choices: h, s
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
