package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	//cards.saveToFile("/Users/mparyani/learninggo.txt")
}
