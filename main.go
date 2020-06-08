package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"

	// In this type of declaration we are telling the Go
	// compiler to figure out the type of the variable.
	// Also we use this type of declaration(:=) only if we
	// are decalring a new variable otherwise just use ("=")
	card := "Ace of Spades"
	card = "Five of Diamonds"

	fmt.Println(card)
}
