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

	// this is for only a single player
	p1Hand := newHand(player1)
	dHand := newHand(dealer)

	// this should be named smth like 'hands'
	// but this is fine for single-player
	players := []*hand{p1Hand}

	// initial deal phase
	for i := 0; i < initialDealSize; i++ {
		for _, p := range players {
			d = deal(d, p)
		}
		d = deal(d, dHand)
	}

	// each player plays until they stand or bust
	for _, p := range players {
		fmt.Println(dHand)

		var done bool
		for !done {
			fmt.Println(p)

			// prompt player to hit or stand
			fmt.Print("(h)it or (s)tand: ")
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
	time.Sleep(time.Second * 1)

	// show dealer's hidden card and total value
	fmt.Printf("dealer's face-down card: '%v'\n", dHand.cards[0])
	fmt.Printf("dealer's cards value: %v\n", dHand.value())

	// determine and announce winner
	dVal := dHand.value()
	for _, p := range players {
		val := p.value()
		if p.bust || val < dVal {
			fmt.Println(p.owner, "lost!")
		} else if val > dVal {
			fmt.Println(p.owner, "won!")
		} else {
			fmt.Println(p.owner, "ties! (push)")
		}
	}
}
