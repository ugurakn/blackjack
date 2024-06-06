package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

const initialDealSize = 2
const initPurseSize = 1000

func main() {
	var numOfPlayers int
	flag.IntVar(&numOfPlayers, "players", 2, "number of players")
	var pNames string
	flag.StringVar(&pNames, "names", "", "enter player names with commas (and no spaces) in between, e.g., name1,name2,name3")
	var decks int
	flag.IntVar(&decks, "decks", 2, "number of decks to use (min 1). The shoe will be auto-reshuffled when there are too few cards left.")

	flag.Parse()

	players := getPlayers(numOfPlayers, pNames)
	sh := newShoe(decks)

	// main loop
	for {
		time.Sleep(time.Second * 1)
		fmt.Println()

		// reshuffle if too few cards remained
		if len(sh.cards)*2 < sh.initSize {
			sh.reshuffle()
			fmt.Println("reshuffled the cards...")
		}

		hands := getHands(players)
		if len(hands) == 0 {
			fmt.Println("---No players left to play---")
			break
		}
		dHand := newHand(newDealer())

		// initial deal phase
		for i := 0; i < initialDealSize; i++ {
			for _, h := range hands {
				deal(sh, h)
			}
			deal(sh, dHand)
		}

		// show dealt cards for each player
		fmt.Println("---Cards dealt---")
		fmt.Println(dHand)

		for _, h := range hands {
			fmt.Println(h)
		}
		time.Sleep(1 * time.Second)

		// check natural blackjacks
		for _, h := range hands {
			if checkBJ(h) {
				fmt.Println(h.owner, "has a blackjack!")
			}
		}

		// each player plays until they stand or bust
		for _, h := range hands {
			if h.bjack {
				continue
			}

			fmt.Println()
			fmt.Println(dHand)
			playTurn(sh, h)
		}

		fmt.Println()
		fmt.Println("---All players have played, dealer's turn---")
		time.Sleep(time.Millisecond * 500)

		// show dealer's hidden card and total value, dealer's turn
		fmt.Printf("dealer's face-down card: '%v'\n", dHand.cards[0])
		if checkBJ(dHand) {
			fmt.Println("dealer has a blackjack!")
		} else {
			fmt.Printf("dealer's cards value: %v\n", dHand.value())
			time.Sleep(time.Second * 1)
			playDealer(sh, dHand)
		}

		fmt.Println("---Game ended---")
		time.Sleep(time.Second * 1)

		// determine win/lose states and payouts for bets
		// then display them
		displayGameEnd(hands, dHand)

		fmt.Println()
		fmt.Print("(q)uit or continue: ")
		quit := getInput()
		if strings.ToLower(quit) == "q" {
			break
		}
	}
}
