package deck

type DeckInterface interface {
	AddCards(cards []string)
	RemoverCard(index int) string
}

// Represents the deck of cards.
type Deck struct {
	Cards []string
}

func (d *Deck) AddCards(cards []string) {
	d.Cards = append(d.Cards, cards...)
}

func (d *Deck) RemoverCard(index int) string {
	card := d.Cards[index]
	d.Cards = append(d.Cards[:index], d.Cards[index+1:]...)
	return card
}
