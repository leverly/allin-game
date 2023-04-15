package main

import "math/rand"

var cardSuits = []string{
	"Spades", "Diamonds", "Clubs", "Hearts",
}

var CardValues = []string{
	"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A",
}

type Card struct {
	Suit  string
	Value string
}

type Cards struct {
	cards []Card
	index int
}

func (c *Cards) Init() {
	for _, suit := range cardSuits {
		for _, value := range CardValues {
			card := Card{Suit: suit, Value: value}
			c.cards = append(c.cards, card)
		}
	}
	c.index = 0
}

func (c *Cards) Shuffle() {
	rand.Shuffle(len(c.cards), func(i, j int) {
		c.cards[i], c.cards[j] = c.cards[j], c.cards[i]
	})
}

func (c *Cards) Next() Card {
	// panic(len(c.cards) <= c.index)
	card := c.cards[c.index]
	c.index++
	return card
}
