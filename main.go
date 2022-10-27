package main

import "fmt"

func main() {
	cards := newDeckFromFile("/Users/mparyani/learninggo.txt")
	fmt.Println(cards)
	//cards.saveToFile("/Users/mparyani/learninggo.txt")
}
