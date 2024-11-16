package player_test

import (
	"game-card/mocks"
	"game-card/player"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

var gameFinished = make(chan bool, 1)

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

	mockPlayer.RemoveCard(strconv.Itoa(6))

	require.Equal(t, []string{"3♥", "2♥", "3♣"}, mockPlayer.Hand)
}

func TestGetCardToHand(t *testing.T) {
	mockPlay := setupMockAndPlayer()

	mockPlay.GetCardToHand("3♠")

	require.Equal(t, []string{"3♥", "2♥", "6♦", "3♣", "3♠"}, mockPlay.Hand)
}

func TestPlayHasFourEqualCards(t *testing.T) {

	testsCards := []struct {
		name           string
		hand           []string
		expectedResult bool
	}{
		{
			name:           "Hand Four cards of a kind.",
			hand:           []string{"3♥", "3♥", "3♦", "3♣"},
			expectedResult: true,
		},
		{
			name:           "Hand with three cards of a kind.",
			hand:           []string{"5♥", "6♥", "5♦", "5♣"},
			expectedResult: false,
		},
		{
			name:           "Hand with all different cards.",
			hand:           []string{"3♥", "6♥", "1♦", "4♣"},
			expectedResult: false,
		},
	}

	for _, tc := range testsCards {
		t.Run(tc.name, func(t *testing.T) {
			mockPlayer := &mocks.MockPlayer{
				Player: player.Player{
					Hand: tc.hand,
				},
			}

			result := mockPlayer.HasFourEqualCards()
			require.Equal(t, tc.expectedResult, result, "Expected result '%v' for a hand %v.", tc.expectedResult, tc.hand)
		})
	}
}

func TestPlay(t *testing.T) {
	leftPile := mocks.NewMockPile(2)
	rightPile := mocks.NewMockPile(2)

	rightPile.Push("3♦")
	rightPile.Push("3♠")

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

	go mockPlayer.Play(&wg, gameFinished)

	wg.Wait()

	require.True(t, mockPlayer.Won, "Expected player victory.")
	require.ElementsMatch(t, mockPlayer.Hand, []string{"3♥", "3♦", "3♠", "3♣"})

}
