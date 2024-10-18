package utils_test

import (
	//deck "game-card/Deck"
	// player "game-card/Player"
	utils "game-card/Utils"
	"testing"

	"github.com/stretchr/testify/require"
)

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

type MockPlayer struct {
	Hand []string
}

func TestGenerateCardsForTheDeck(t *testing.T) {
	ranks := []string{"1", "2", "3"}
	suits := []string{"♥", "♦", "♣", "♠"}
	deckHand := []string{"1♥", "1♦", "1♣", "1♠", "2♥", "2♦", "2♣", "2♠", "3♥", "3♦", "3♣", "3♠"}

	mockDeck := &MockDeck{}
	utils.GenerateCardsForTheDeck(mockDeck, ranks, suits)

	require.Equal(t, mockDeck.Cards, deckHand)
}

// TODO: REFAZER
// func TestCardsPartitionForPlayer(t *testing.T) {
// 	deckHand := deck.Deck{}
// 	deckHand.Cards = []string{"1♥", "1♦", "1♣", "1♠", "2♥", "2♦", "2♣", "2♠", "3♥", "3♦", "3♣", "3♠"}

//		require.Equal(t, utils.DealCardsToPlayers(&deckHand, &player.Player{}), []string{"1♦"})
//	}
//func TestDealCardsToPlayer(t *testing.T) {}
