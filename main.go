package main

import (
	"fmt"
	"game-card/deck"
	"game-card/pile"
	"game-card/player"
	"game-card/utils"
	"sync"
)

var (
	suits        []string = []string{"♥", "♦", "♣", "♠"}
	ranks        []string = []string{"1", "2", "3", "4", "5", "6"}
	wg           sync.WaitGroup
	gameFinished = make(chan bool, 1)
)

func main() {

	// Generating the Deck
	deckCards := deck.Deck{}

	utils.GenerateCardsForTheDeck(&deckCards, ranks, suits)

	// Creating the Stacks
	pileOne := pile.NewPile("Ana-Beto", 3)
	pileTwo := pile.NewPile("Beto-Carla", 3)
	pileThree := pile.NewPile("Carla-Daniel", 3)
	pileFour := pile.NewPile("Daniel-Ana", 3)

	// Generating cards for the stacks
	utils.DealCardsToPile(&deckCards, pileOne)
	utils.DealCardsToPile(&deckCards, pileTwo)
	utils.DealCardsToPile(&deckCards, pileThree)
	utils.DealCardsToPile(&deckCards, pileFour)

	// Creating the players
	players := []*player.Player{}
	playerAna := player.NewPlayer("Ana", []string{}, pileOne, pileFour, 0)
	playerBeto := player.NewPlayer("Beto", []string{}, pileTwo, pileOne, 0)
	playerCarla := player.NewPlayer("Carla", []string{}, pileThree, pileTwo, 0)
	playerDaniel := player.NewPlayer("Daniel", []string{}, pileFour, pileThree, 0)

	// Generating cards for players
	utils.DealCardsToPlayer(&deckCards, playerAna)
	utils.DealCardsToPlayer(&deckCards, playerBeto)
	utils.DealCardsToPlayer(&deckCards, playerCarla)
	utils.DealCardsToPlayer(&deckCards, playerDaniel)

	// Printing the stacks
	fmt.Println("********************** Piles **************************************")
	fmt.Println(pileOne)
	fmt.Println(pileTwo)
	fmt.Println(pileThree)
	fmt.Println(pileFour)
	fmt.Println("***********************************************************************")

	// Printing the players
	fmt.Println("********************** Players **************************************")
	fmt.Println(playerAna)
	fmt.Println(playerBeto)
	fmt.Println(playerCarla)
	fmt.Println(playerDaniel)
	fmt.Println("***********************************************************************")

	players = append(players, playerAna, playerBeto, playerCarla, playerDaniel)

	for _, player := range players {
		wg.Add(1)
		go player.Play(&wg, gameFinished)
	}

	<-gameFinished
	close(gameFinished)
	wg.Wait()

}
