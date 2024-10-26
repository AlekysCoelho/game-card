package mocks

type MockPlayer struct {
	Hand []string
}

func (m *MockPlayer) GetCardToHand(card string) {
	m.Hand = append(m.Hand, card)
}

func (p *MockPlayer) HasFourEqualCards() bool {
	numbersFrequency := make(map[string]int)

	for _, card := range p.Hand {
		numbersFrequency[string(card[0])]++
	}

	return len(numbersFrequency) == 1
}

func (p *MockPlayer) HasCard() bool {
	return len(p.Hand) >= 1
}
