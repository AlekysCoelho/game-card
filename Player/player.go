package player

import pile "game-card/Pile"

type PlayerInterface interface {
	GetCard(card string)
	HasFourHand() bool
}

type Player struct {
	Name         string
	Hand         []string
	LeftPile     *pile.Pile
	RightPile    *pile.Pile
	NumberOfPlay int
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
func (p *Player) GetCard(card string) {
	p.Hand = append(p.Hand, card)
}

// Checks if the player has a Four hand.
func (p *Player) HasFourHand() bool {
	numbersFrequency := make(map[string]int)

	for _, card := range p.Hand {
		numbersFrequency[string(card[0])]++
	}

	if len(numbersFrequency) == 1 {
		return true
	}

	return false
}
