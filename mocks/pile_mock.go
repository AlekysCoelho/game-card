package mocks

type MockPile struct {
	Name  string
	Cards []string
	ch    chan string
}

func NewMockPile(size int) *MockPile {
	return &MockPile{
		Cards: make([]string, 0, size),
		ch:    make(chan string, size),
	}
}

func (m *MockPile) Push(card string) {

	if len(m.Cards) < cap(m.Cards) {
		m.Cards = append(m.Cards, card)
		m.ch <- card
	}
}

func (m *MockPile) Pop() string {

	if len(m.Cards) == 0 {
		return ""
	}
	card := m.Cards[len(m.Cards)-1]
	m.Cards = m.Cards[:len(m.Cards)-1]

	return card
}

func (m *MockPile) HasCard() bool {
	return len(m.Cards) >= 1
}

// Função para ver a capacidade de `Cards` e `ch`
func (p *MockPile) Capacity() (int, int) {
	return cap(p.Cards), cap(p.ch)
}
