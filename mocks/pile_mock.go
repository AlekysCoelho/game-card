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

	select {
	case m.ch <- card:
	default:
	}
}

func (m *MockPile) Pop() string {

	select {
	case card := <-m.ch:

		return card
	default:
		return ""
	}
}

func (m *MockPile) HasCard() bool {
	return len(m.Cards) >= 1
}

// Função para ver a capacidade de `Cards` e `ch`
func (m *MockPile) Capacity() (int, int) {
	return len(m.ch), cap(m.ch)
}

func (m *MockPile) GetName() string {
	return m.Name
}
