package game_test

import (
	"game-card/game"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReturnsTheLowestFrequencyAmongCards(t *testing.T) {
	tests := []struct {
		name         string
		cards        []string
		expectedCard string
	}{
		{"Test to rule out 1", []string{"1♥", "2♥", "2♦", "2♣"}, "1"},
		{"Test to rule out 2", []string{"1♥", "1♦", "1♣", "2♥"}, "2"},
		{"Test to rule out 3", []string{"3♥", "3♣", "6♣", "6♦"}, "3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minFrequency := game.ReturnsTheLowestFrequencyAmongCards(tt.cards)

			require.Equal(t, tt.expectedCard, minFrequency, "Expected: %s but got: %s", tt.expectedCard, minFrequency)
		})
	}
}

func TestDiscardTheCardWithLowestFrequency(t *testing.T) {
	tests := []struct {
		name         string
		cards        []string
		expectedCard string
		minFrequency string
	}{
		{"Discard card with frequency '1': expected '1♥' ", []string{"1♥", "2♥", "2♦", "2♣"}, "1♥", "1"},
		{"Discard card with frequency '2': expected '2♥' ", []string{"1♥", "1♦", "1♣", "2♥"}, "2♥", "2"},
		{"Discard card with frequency '3': expected '3♥' ", []string{"3♥", "3♣", "6♣", "6♦"}, "3♥", "3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cardDiscard := game.DiscardTheCardWithLowestFrequency(tt.cards, tt.minFrequency)

			require.Equal(t, tt.expectedCard, cardDiscard, "Expected: %s but got: %s", tt.expectedCard, cardDiscard)
		})
	}
}
