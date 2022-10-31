package main

import "fmt"

type bot interface {
	getGretting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGretting(eb)
	printGretting(sb)

}

func (eb englishBot) getGretting() string {
	return "hello there"
}

func (sb spanishBot) getGretting() string {
	return "hola"
}

func printGretting(b bot) {
	fmt.Println(b.getGretting())
}
