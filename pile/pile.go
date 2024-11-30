package pile

import (
	logs "game-card/Logs"
)

type PileInterface interface {
	Push(card string)
	Pop() string
	Capacity() (int, int)
	GetName() string
}

// Represents the pile of cards.
// Cards: stores the cards.
// ch: channel for manipulating cards in the pile.
type Pile struct {
	Name string
	ch   chan string
}

// New Pile instance.
func NewPile(name string, size int) *Pile {
	return &Pile{
		Name: name,
		ch:   make(chan string, size),
	}
}

func (p *Pile) Push(card string) {
	select {
	case p.ch <- card:
		logs.Log.Info("Stack Push", "Stack", p.Name, "Added card", card)
	default:
		logs.Log.Warn("Stack Full", "Stack", p.Name, "Cannot add card", card)
	}
}

// Method for removing a card from the pile.
// Remove the first card; FIFO
func (p *Pile) Pop() string {
	select {
	case card := <-p.ch:
		logs.Log.Info("Stack Pop", "Stack", p.Name, "Removed card", card)
		return card
	default:
		logs.Log.Warn("Stack Empty", "Stack", p.Name)
		return ""
	}
}

// Function that returns the capacity and size of the channel
func (p *Pile) Capacity() (int, int) {
	return len(p.ch), cap(p.ch)
}

func (p *Pile) GetName() string {
	return p.Name
}
