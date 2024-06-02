package main

import (
	"fmt"
	"time"

	"github.com/ugurakn/deck"
)

const initialDealSize = 2

func main() {
	sh := new(shoe)
	sh.cards = deck.New(deck.Shuffle)
	sh.initSize = len(sh.cards)

	dHand := newHand(dealer, false)

	// this should be named smth like 'hands'
	// but this is fine for single-hand per player
	players := []*hand{newHand(player1, true), newHand(player2, true)}

	// initial deal phase
	for i := 0; i < initialDealSize; i++ {
		for _, p := range players {
			deal(sh, p)
		}
		deal(sh, dHand)
	}

	// show dealt cards for each player
	fmt.Println("All cards dealt...")
	fmt.Println(dHand)
	for _, p := range players {
		fmt.Println(p)
	}
	time.Sleep(1 * time.Second)

	// check natural blackjacks
	// if dealer has one, skip players turn
	dealerbjack := checkBJ(dHand)
	if dealerbjack {
		fmt.Println("dealer has a blackjack!")
	}

	for _, p := range players {
		if checkBJ(p) {
			fmt.Println(p.owner, "has a blackjack!")
		}
	}

	// each player plays until they stand or bust
	for _, p := range players {
		// skip players turn if dealer has a bjack
		if dealerbjack {
			break
		}

		// skip this player's turn if he has a bjack
		if p.bjack {
			continue
		}

		fmt.Println()
		fmt.Println(dHand)
		playTurn(sh, p)
	}
	fmt.Println("Game ended...")
	time.Sleep(time.Millisecond * 500)

	// show dealer's hidden card and total value
	fmt.Printf("dealer's face-down card: '%v'\n", dHand.cards[0])
	fmt.Printf("dealer's cards value: %v\n", dHand.value())
	time.Sleep(time.Millisecond * 750)

	// determine and announce winner
	for _, p := range players {
		p.setWinState(dHand)
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
