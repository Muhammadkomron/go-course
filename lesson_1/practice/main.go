package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards)
	cards.shuffle()
	fmt.Println(cards)
	cards.saveToFile("qwerty")
	getCards := newDeckFromFile("qwerty")
	getCards.print()
}
