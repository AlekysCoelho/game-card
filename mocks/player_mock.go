package mocks

import (
	"game-card/player"
	"sync"
)

type MockPlayer struct {
	player.Player
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

func (mp *MockPlayer) Play(wg *sync.WaitGroup) {
	mp.Player.Play(wg)
}
