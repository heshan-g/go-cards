package main
 
func main() {
	cards := newDeck()

	// cards.print()
	// fmt.Println("------------------------")

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	// fmt.Println("------------------------")

	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards.shuffle()
	cards.print()
}
