package player

import (
	logs "game-card/Logs"
	game "game-card/game"
	pile "game-card/pile"
	"sync"
	"time"
)

var (
	mu sync.Mutex
)

type PlayerInterface interface {
	GetCardToHand(card string)
	HasFourEqualCards() bool
	HasCard() bool
	RemoveCard(cardLowerFrequency string)
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

func (p *Player) RemoveCard(cardLowerFrequency string) {

	for index, card := range p.Hand {
		if string(card[0]) == cardLowerFrequency {
			p.Hand = append(p.Hand[:index], p.Hand[index+1:]...)
			logs.Log.Info("Player RemoveCard", "Player", p.Name, "Hand", p.Hand, "Carta Removida", card)
			p.LeftPile.Push(card)
			return
		}
	}
}

// Player takes a card and puts it in his hand.
func (p *Player) GetCardToHand(card string) {
	logs.Log.Info("Player GetCard", "Player", p.Name, "Hand", p.Hand, "Pegou a carta", card)
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
func (p *Player) Play(wg *sync.WaitGroup, gameFinished chan bool) {
	defer wg.Done()
	logs.Log.Info("Começou a jogar", "Player", p.Name, "Hand", p.Hand)

	for !p.Won {

		select {
		case <-gameFinished:
			// The game is over
			logs.Log.Info("Game Finished", "Player stops playing as game is over.", p.Name)
			return
		default:

			lenPile, capPile := p.LeftPile.Capacity()

			if lenPile < capPile {
				cardLowerFrequency := game.ReturnsTheLowestFrequencyAmongCards(p.Hand)
				p.RemoveCard(cardLowerFrequency)
				if len(p.Hand) < 4 {

					newCard := p.RightPile.Pop()
					if len(newCard) == 0 {
						return
					}
					p.Hand = append(p.Hand, newCard)
					logs.Log.Info("Pegou uma carta", "Player", p.Name, "Hand", p.Hand, "Carta pegue", newCard)
				}
			}

			if p.HasFourEqualCards() {
				p.Won = true
				logs.Log.Info("SUCCESS -> We have a winner!", "Player", p.Name, "Cards", p.Hand)

				mu.Lock()
				gameFinished <- true
				mu.Unlock()

				return
			}

			time.Sleep(time.Microsecond * 100)
		}

	}
}
