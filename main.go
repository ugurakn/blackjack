package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ugurakn/deck"
)

const initialDealSize = 2

func main() {
	d := deck.New(deck.Shuffle)

	dHand := newHand(dealer)

	// this should be named smth like 'hands'
	// but this is fine for single-player
	players := []*hand{newHand(player1), newHand(player2)}

	// initial deal phase
	for i := 0; i < initialDealSize; i++ {
		for _, p := range players {
			d = deal(d, p)
		}
		d = deal(d, dHand)
	}

	// check natural blackjacks, skip game loop if any exists
	var blackjack bool
	blackjack = checkBJ(dHand)
	if blackjack {
		fmt.Println("dealer has a blackjack!")
	}
	for _, p := range players {
		pBj := checkBJ(p)
		if pBj {
			blackjack = true
			fmt.Println(p.owner, "has a blackjack!")
		}
	}

	// each player plays until they stand or bust
	for _, p := range players {
		if blackjack {
			break
		}

		fmt.Println()
		fmt.Println(dHand)

		var done bool
		for !done {
			fmt.Println(p)

			// prompt player to hit or stand
			fmt.Printf("(%v) (h)it or (s)tand: ", p.owner)
			var in string
			fmt.Scanf("%s\n", &in)
			switch strings.ToLower(in) {
			case "h":
				d = deal(d, p)
				fmt.Printf("%v HIT: got %v. new value:%v\n", p.owner, p.cards[len(p.cards)-1], p.value())
				time.Sleep(time.Second * 1)
				// check bust
				if val := p.value(); val > 21 {
					done = true
					fmt.Printf("%v BUST!\n", p.owner)
				}

			case "s":
				fmt.Println(p.owner, "STAND")
				done = true

			default:
				fmt.Println("unknown input:", in)
				os.Exit(1)
			}
		}
	}
	fmt.Println("All players played. Dealer's turn...")
	time.Sleep(time.Millisecond * 500)

	// show dealer's hidden card and total value
	fmt.Printf("dealer's face-down card: '%v'\n", dHand.cards[0])
	fmt.Printf("dealer's cards value: %v\n", dHand.value())

	// determine and announce winner
	if blackjack {
		if dHand.bjack {
			for _, p := range players {
				if p.bjack {
					p.winState = push
				}
				p.winState = lost
			}
		} else {
			for _, p := range players {
				if p.bjack {
					p.winState = winbj
				}
			}
		}
	}

	for _, p := range players {
		if p.winState == undecided {
			p.setWinState(dHand)
		}
	}

	// display winState messages
	for _, p := range players {
		switch p.winState {
		case lost:
			fmt.Println(p.owner, "LOST!")
		case bust:
			fmt.Println(p.owner, "LOST! (bust)")
		case push:
			fmt.Println(p.owner, "PUSH!")
		case win:
			fmt.Println(p.owner, "WON!")
		case winbj:
			fmt.Println(p.owner, "WON! (blackjack)")
		case undecided:
			panic("player winState shouldn't be undecided")
		}
	}
}
