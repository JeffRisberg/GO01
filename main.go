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

// Create a new type called 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
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

	deck := newDeck()
	card := newCard()
	cardValue := newCardValue()

	fmt.Println(deck)
	fmt.Println(card)
	fmt.Println(cardValue)

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func newCardValue() int {
	return 1
}
