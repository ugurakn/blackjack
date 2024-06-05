package main

import (
	"fmt"
	"time"

	"github.com/ugurakn/deck"
)

const initialDealSize = 2
const initPurseSize = 1000

func main() {
	sh := new(shoe)
	sh.cards = deck.New(deck.Shuffle)
	sh.initSize = len(sh.cards)

	// create dealer & players
	p1 := &player{name: "player1", purse: initPurseSize, isDealer: false, isHuman: true}
	p2 := &player{name: "player2", purse: initPurseSize, isDealer: false, isHuman: true}

	dHand := newHand(newDealer())

	// represents players' hands
	hands := []*hand{newHand(p1), newHand(p2)}

	// initial deal phase
	for i := 0; i < initialDealSize; i++ {
		for _, h := range hands {
			deal(sh, h)
		}
		deal(sh, dHand)
	}

	// show dealt cards for each player
	fmt.Println("All cards dealt...")
	fmt.Println(dHand)

	for _, h := range hands {
		fmt.Println(h)
	}
	time.Sleep(1 * time.Second)

	// check natural blackjacks
	// dealerbjack := checkBJ(dHand)

	for _, h := range hands {
		if checkBJ(h) {
			fmt.Println(h.owner, "has a blackjack!")
		}
	}

	// each player plays until they stand or bust
	for _, h := range hands {
		// skip this hand's turn if hand has a bjack
		if h.bjack {
			continue
		}

		fmt.Println()
		fmt.Println(dHand)
		playTurn(sh, h)
	}
	fmt.Println("---All players have played, dealer's turn---")
	time.Sleep(time.Millisecond * 500)

	// show dealer's hidden card and total value
	fmt.Printf("dealer's face-down card: '%v'\n", dHand.cards[0])
	fmt.Printf("dealer's cards value: %v\n", dHand.value())
	time.Sleep(time.Second * 1)
	// dealer's turn
	playDealer(sh, dHand)
	fmt.Println("---Game ended---")
	time.Sleep(time.Second * 1)

	// determine win/lose states and payouts for bets
	// then display them
	for _, h := range hands {
		h.setWinState(dHand)
		amount := payout(h)
		switch h.winState {
		case lost:
			fmt.Println(h.owner, "LOST!")
			fmt.Printf("%v purse: %v (%v)\n", h.owner, h.owner.purse, amount)
		case bust:
			fmt.Println(h.owner, "LOST! (bust)")
			fmt.Printf("%v purse: %v (%v)\n", h.owner, h.owner.purse, amount)
		case push:
			fmt.Println(h.owner, "PUSH!")
			fmt.Printf("%v purse: %v (%v)\n", h.owner, h.owner.purse, amount)
		case win:
			fmt.Println(h.owner, "WON!")
			fmt.Printf("%v purse: %v (+%v)\n", h.owner, h.owner.purse, amount)
		case winbj:
			fmt.Println(h.owner, "WON! (blackjack)")
			fmt.Printf("%v purse: %v (+%v)\n", h.owner, h.owner.purse, amount)
		case undecided:
			panic("player winState shouldn't be undecided")
		}
	}
}
