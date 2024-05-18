package main

func main() {

	// create a new deck
	cards0 := newDeck()

	// deal from the give deck of a given handSize
	hand, remainingCards := deal(cards0, 5)

	// print the hand and remianing cards
	hand.print()
	remainingCards.print()

	// create a new deck and save it to file
	cards1 := newDeck()
	cards1.saveToFile("my_cards")

	// read values from file and print the same
	cards2 := newDeckFromFile("my_cards")
	cards2.print()

	// create a new deck and shuffle and print the same
	cards3 := newDeck()
	cards3.shuffle()
	cards3.print()

}
