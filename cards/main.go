package main

func main() {
	var cards deck = newDeck()
	cards.shuffle()
	cards.print()

	//hand, remainingCards := deal(cards, 5)

	//fmt.Println(cards.saveToFile("test.txt"))

	//var cards deck = newDeckFromFile("test.txt")
	//cards.print()

}
