package player

import pile "game-card/pile"

type PlayerInterface interface {
	GetCardToHand(card string)
	HasFourHand() bool
	HasCard() bool
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
func (p *Player) GetCardToHand(card string) {
	p.Hand = append(p.Hand, card)
}

// Checks if the player has a Four hand.
func (p *Player) HasFourHand() bool {
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
