package player_test

import (
	"game-card/mocks"
	"game-card/player"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func setupMockAndPlayer() *mocks.MockPlayer {

	leftPile := mocks.NewMockPile(2)
	rightPile := mocks.NewMockPile(2)

	mockPlayer := &mocks.MockPlayer{
		Player: player.Player{
			Name:         "Player 1",
			Hand:         []string{"3♥", "2♥", "6♦", "3♣"},
			LeftPile:     leftPile,
			RightPile:    rightPile,
			NumberOfPlay: 0,
			Won:          false,
		},
	}

	return mockPlayer
}

func TestRemoveCard(t *testing.T) {

	mockPlayer := setupMockAndPlayer()

	card := mockPlayer.RemoveCard(strconv.Itoa(6))

	require.Equal(t, "6♦", card)
	require.Equal(t, []string{"3♥", "2♥", "3♣"}, mockPlayer.Hand)
}

func TestGetCardToHand(t *testing.T) {
	mockPlay := setupMockAndPlayer()

	mockPlay.GetCardToHand("3♠")

	require.Equal(t, []string{"3♥", "2♥", "6♦", "3♣", "3♠"}, mockPlay.Hand)
}

func TestPlay(t *testing.T) {
	leftPile := mocks.NewMockPile(2)
	rightPile := mocks.NewMockPile(2)

	rightPile.Push("3♦")
	rightPile.Push("3♠")
	// rightPile.Push("1♠")

	mockPlayer := &mocks.MockPlayer{
		Player: player.Player{
			Name:         "Player 1",
			Hand:         []string{"3♥", "2♥", "6♦", "3♣"},
			LeftPile:     leftPile,
			RightPile:    rightPile,
			NumberOfPlay: 0,
			Won:          false,
		},
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go mockPlayer.Play(&wg)

	wg.Wait()

	require.True(t, mockPlayer.Won, "Expected player victory.")
	require.ElementsMatch(t, mockPlayer.Hand, []string{"3♥", "3♦", "3♠", "3♣"})

}
