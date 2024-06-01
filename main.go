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

	// each player plays until they stand or bust
	for _, p := range players {
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
					p.bust = true
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
	// possible cases: lost, lost (bust), push (i.e. draw), won, won (blackjack)

	dVal := dHand.value()
	for _, p := range players {
		val := p.value()
		if p.bust {
			fmt.Println(p.owner, "LOST! (bust)")
		} else if val < dVal {
			fmt.Println(p.owner, "LOST!")
		} else if val > dVal {
			fmt.Println(p.owner, "WON!")
		} else {
			fmt.Println(p.owner, "PUSH!")
		}
	}
}
