package mocks

import (
	"game-card/player"
	"sync"
)

type MockPlayer struct {
	player.Player
}

func (p *MockPlayer) RemoveCard(cardLowerFrequency string) {

	for index, card := range p.Hand {
		if string(card[0]) == cardLowerFrequency {
			p.Hand = append(p.Hand[:index], p.Hand[index+1:]...)
			p.LeftPile.Push(card)
			return
		}
	}
}
func (mp *MockPlayer) GetCardToHand(card string) {
	mp.Hand = append(mp.Hand, card)
}

func (mp *MockPlayer) HasFourEqualCards() bool {
	numbersFrequency := make(map[string]int)

	for _, card := range mp.Hand {
		numbersFrequency[string(card[0])]++
	}

	return len(numbersFrequency) == 1
}

func (mp *MockPlayer) HasCard() bool {
	return len(mp.Hand) >= 1
}

func (mp *MockPlayer) Play(wg *sync.WaitGroup, gameFinished chan bool) {
	mp.Player.Play(wg, gameFinished)
}
