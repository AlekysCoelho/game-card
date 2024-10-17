package player

import pile "game-card/Pile"

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
