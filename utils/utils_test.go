package utils_test

import (
	"fmt"
	deckMock "game-card/mocks"
	deckPile "game-card/mocks"
	deckPlayer "game-card/mocks"
	utils "game-card/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateCardsForTheDeck(t *testing.T) {
	ranks := []string{"1", "2", "3"}
	suits := []string{"♥", "♦", "♣", "♠"}
	deckHand := []string{"1♥", "1♦", "1♣", "1♠", "2♥", "2♦", "2♣", "2♠", "3♥", "3♦", "3♣", "3♠"}

	mockDeck := &deckMock.MockDeck{}
	utils.GenerateCardsForTheDeck(mockDeck, ranks, suits)

	require.Equal(t, mockDeck.Cards, deckHand)
}

func TestDealCardsToPlayer(t *testing.T) {

	deck := []string{"1♥", "1♦", "1♣", "1♠", "2♥", "2♦", "2♣", "2♠", "3♥", "3♦", "3♣", "3♠"}

	mockDeck := &deckMock.MockDeck{}
	mockDeck.AddCards(deck)
	mockPlayer := &deckPlayer.MockPlayer{}

	err := utils.DealCardsToPlayer(mockDeck, mockPlayer)
	if err != nil {
		fmt.Println(err)
	}

	require.Equal(t, 4, len(mockPlayer.Hand))
}

func TestDealCardsToPile(t *testing.T) {
	deck := []string{"1♥", "1♦", "1♣", "1♠", "2♥", "2♦", "2♣", "2♠", "3♥", "3♦", "3♣", "3♠"}

	mockDeck := &deckMock.MockDeck{}
	mockDeck.AddCards(deck)
	mockPile := deckPile.NewMockPile(2)

	err := utils.DealCardsToPile(mockDeck, mockPile)
	require.NoError(t, err)

	require.Equal(t, 2, len(mockPile.Cards))
}
