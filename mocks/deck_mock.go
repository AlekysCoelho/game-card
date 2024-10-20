package mocks

type MockDeck struct {
	Cards []string
}

func (m *MockDeck) AddCards(cards []string) {
	m.Cards = append(m.Cards, cards...)
}

func (m *MockDeck) RemoverCard(index int) string {
	card := m.Cards[index]
	m.Cards = append(m.Cards[:index], m.Cards[index+1:]...)
	return card
}

func (d *MockDeck) GetCards() []string {
	return d.Cards
}
