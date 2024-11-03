package player

import (
	logs "game-card/Logs"
	game "game-card/game"
	pile "game-card/pile"
	"sync"
	"time"
)

type PlayerInterface interface {
	GetCardToHand(card string)
	HasFourEqualCards() bool
	HasCard() bool
}

type Player struct {
	Name         string
	Hand         []string
	LeftPile     pile.PileInterface
	RightPile    pile.PileInterface
	NumberOfPlay int
	Won          bool
}

func NewPlayer(name string, hand []string, leftPile, rightPile *pile.Pile, numberOfPlay int) *Player {
	return &Player{
		Name:         name,
		Hand:         hand,
		LeftPile:     leftPile,
		RightPile:    rightPile,
		NumberOfPlay: numberOfPlay,
	}
}

// Player takes a card and puts it in his hand.
func (p *Player) GetCardToHand(card string) {
	p.Hand = append(p.Hand, card)
}

// Checks if the player has a Four hand.
func (p *Player) HasFourEqualCards() bool {
	numbersFrequency := make(map[string]int)

	for _, card := range p.Hand {
		numbersFrequency[string(card[0])]++
	}

	return len(numbersFrequency) == 1
}

// Checks if the hand has cards.
func (p *Player) HasCard() bool {
	return len(p.Hand) >= 1
}

// Function for the player to play.
// Discard a card from the left pile and take a card from the right pile.
// The player will win if he makes a hand Four.
func (p *Player) Play(wg *sync.WaitGroup) {
	defer wg.Done()

	for !p.Won {

		indiceCard := game.ReturnsTheLowestFrequencyAmongCards(p.Hand)
		cardToDiscard := game.DiscardTheCardWithLowestFrequency(p.Hand, indiceCard)
		p.LeftPile.Push(cardToDiscard)

		newCard := p.RightPile.Pop()
		p.Hand = append(p.Hand, newCard)

		if p.HasFourEqualCards() {
			p.Won = true
			logs.Log.Info("We have a winner!", "Player", p.Name, "Cards", p.Hand)
		}

		time.Sleep(time.Microsecond * 100)
	}
}
