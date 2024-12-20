package game

// Count the frequency of equal cards.
// Returns the lowest frequency among the cards.
func ReturnsTheLowestFrequencyAmongCards(cards []string) (cardLowerFrequency string) {
	numbersFrequency := make(map[string]int)

	for _, card := range cards {
		numbersFrequency[string(card[0])]++
	}

	minFrequency := -1

	for card, count := range numbersFrequency {
		if minFrequency == -1 || count < minFrequency {
			minFrequency = count
			cardLowerFrequency = card
		}
	}
	// fmt.Println("TYPE -> ", reflect.TypeOf(lowersCard))

	return
}

// Discard the card less frequently
// cards: Hand player
// minFrequency: minimum frequency returned from function ReturnsTheLowestFrequencyAmongCards
func DiscardTheCardWithLowestFrequency(cards []string, minFrequency string) string {
	for _, card := range cards {
		if string(card[0]) == minFrequency {

			return card
		}
	}
	return ""
}
