package card

import "fmt"

func PrintCards(cards []Card) string {
	var value string
	for _, card := range cards {
		value += fmt.Sprintf("(%s, %s)", card.suit, card.value.stringValue)
	}
	return value
}

func Combinations(input []Card, n int) [][]Card {
	if n == 0 {
		return [][]Card{{}}
	}
	if len(input) < n {
		return [][]Card{}
	}
	first := input[0]
	rest := input[1:]
	combs1 := Combinations(rest, n-1)
	combs2 := Combinations(rest, n)
	for i := range combs1 {
		combs1[i] = append(combs1[i], first)
	}
	return append(combs1, combs2...)
}
