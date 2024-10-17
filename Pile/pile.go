package pile

import logs "game-card/Logs"

// Represents the pile of cards.
// Cards: stores the cards.
// ch: channel for manipulating cards in the pile.
type Pile struct {
	Name  string
	Cards []string
	ch    chan string
}

// New Pile instance.
func NewPile(name string, size int) *Pile {
	return &Pile{
		Name:  name,
		Cards: make([]string, 0, size),
		ch:    make(chan string, size),
	}
}

// Method for adding card to the pile.
// Add the card at the end.
func (p *Pile) Push(card string) {
	if len(p.Cards) < cap(p.Cards) {
		p.Cards = append(p.Cards, card)
		p.ch <- card
	} else {
		logs.Log.Warn("Full pile", "Pile", p.Name, "Cards", p.Cards)
	}
}

// Method for removing a card from the pile.
// Remove the first card; FIFO
func (p *Pile) Pop() []string {
	if len(p.Cards) == 0 {
		logs.Log.Warn("Empty pile", "Pile", p.Name, "Cards", p.Cards)
		return []string{""}
	}
	card := p.Cards[:1]
	p.Cards = p.Cards[1:]

	return card
}
