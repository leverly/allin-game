package card

import "math/rand"

var cardTypes = []string{
	"Spades", "Diamonds", "Clubs", "Hearts",
}

var cardValues = []CardValue{
	{"2", 2}, {"3", 3}, {"4", 4},
	{"5", 5}, {"6", 6}, {"7", 7},
	{"8", 8}, {"9", 9}, {"10", 10},
	{"J", 11}, {"Q", 12}, {"K", 13},
	{"A", 14},
}

type FullCards struct {
	cards    []Card
	iterator int
}

func (f *FullCards) Init() {
	for _, suit := range cardTypes {
		for _, value := range cardValues {
			card := Card{suit: suit, value: value}
			f.cards = append(f.cards, card)
		}
	}
	f.iterator = 0
}

func (f *FullCards) Shuffle() {
	rand.Shuffle(len(f.cards), func(i, j int) {
		f.cards[i], f.cards[j] = f.cards[j], f.cards[i]
	})
}

// get next card
func (f *FullCards) Next() Card {
	if f.iterator >= len(f.cards) {
		panic("iterator out of range")
	}
	card := f.cards[f.iterator]
	f.iterator++
	return card
}
