package utils_test

import (
	"fmt"
	"game-card/mocks"
	utils "game-card/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func setupMockAndPlayer() (*mocks.MockDeck, *mocks.MockPlayer) {
	deck := []string{"1♥", "1♦", "1♣", "1♠", "2♥", "2♦", "2♣", "2♠", "3♥", "3♦", "3♣", "3♠"}
	mockDeck := &mocks.MockDeck{}
	mockDeck.AddCards(deck)
	mockPlayer := &mocks.MockPlayer{}
	return mockDeck, mockPlayer
}

func setupMockPile(size int) *mocks.MockPile {
	return mocks.NewMockPile(size)
}

func TestGenerateCardsForTheDeck(t *testing.T) {
	ranks := []string{"1", "2", "3"}
	suits := []string{"♥", "♦", "♣", "♠"}
	deckHand := []string{"1♥", "1♦", "1♣", "1♠", "2♥", "2♦", "2♣", "2♠", "3♥", "3♦", "3♣", "3♠"}

	mockDeck := &mocks.MockDeck{}
	utils.GenerateCardsForTheDeck(mockDeck, ranks, suits)

	require.Equal(t, mockDeck.Cards, deckHand)
}

func TestDealCardsToPlayer(t *testing.T) {

	mockDeck, mockPlayer := setupMockAndPlayer()

	err := utils.DealCardsToPlayer(mockDeck, mockPlayer)
	if err != nil {
		fmt.Println(err)
	}

	require.Equal(t, 4, len(mockPlayer.Hand))
}

func TestDealCardsToPile(t *testing.T) {

	mockDeck, _ := setupMockAndPlayer()

	mockPile := setupMockPile(2)

	err := utils.DealCardsToPile(mockDeck, mockPile)
	require.NoError(t, err)

	require.Equal(t, 2, len(mockPile.Cards))
}
