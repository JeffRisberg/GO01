package main

import (
	"fmt"
	"strings"
)

func newCard() string {
	var card string

	card = "Ace of Spades"

	return card
}

func main() {
	var count int16 = 1

	if count == 0 {
		fmt.Println("Hello World")
	} else {
		fmt.Println("Goodbye World")
	}

	greetings := []string{"Hello", "world!"}
	fmt.Println(strings.Join(greetings, " "))

	var sum = 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum) // 10 (1+2+3+4)

	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n) // 8 (1*2*2*2)

	card := newCard()
	cardValue := newCardValue()

	fmt.Println(card)
	fmt.Print(cardValue)
}

func newCardValue() int {
	return 1
}
