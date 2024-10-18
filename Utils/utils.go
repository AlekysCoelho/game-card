package utils

import (
	"errors"
	deck "game-card/Deck"
	player "game-card/Player"
	"math/rand"
	"time"
)

// Function to generate the deck of cards.
// Ranks: Slice of strings with values ​​from 1 to 6
// Suits: Slice of strings with the suits "♥", "♦", "♣", "♠"
// func GenerateCardsForTheDeck(d *deck.Deck, ranks, suits []string) []string {
// 	var card string
// 	for _, r := range ranks {
// 		for _, s := range suits {
// 			card = r + s
// 			d.Cards = append(d.Cards, card)
// 		}
// 	}

// 	return d.Cards
// }

// Function to generate the deck of cards.
// Ranks: Slice of strings with values ​​from 1 to 6
// Suits: Slice of strings with the suits "♥", "♦", "♣", "♠"
func GenerateCardsForTheDeck(deck deck.DeckInterface, ranks, suits []string) {
	cards := []string{}

	for _, rank := range ranks {
		for _, suit := range suits {
			cards = append(cards, rank+suit)
		}
	}

	deck.AddCards(cards)
}

// Dealing the cards to the player. Function to distribute cards to the player
// Deck: deck of cards
// player: Player's hand
func DealCardsToPlayer(d *deck.Deck, player *player.Player) error {
	if len(d.Cards) == 0 {
		return errors.New("deck is empty")
	}
	if len(player.Hand) != 0 {
		return errors.New("your hand already has cards")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i <= 4; i++ {
		index := r.Intn(len(d.Cards))
		card := d.RemoverCard(index)
		player.GetCard(card)
		//card := d.Cards[index]
		//player.Hand = append(player.Hand, card)
		//d.Cards = append(d.Cards[:index], d.Cards[index+1:]...)
	}

	return nil
}

func CountsTheFrequencyOfNumbersInAHandOfCards(cards []string) {}
