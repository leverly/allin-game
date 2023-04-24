package card

import (
	"sort"
)

const (
	CardTypeStraightFlush = 9
	CardTypeIsFour        = 8
	CardTypeIsFullHouse   = 7
	CardTypeIsFlush       = 6
	CardTypeIsStraight    = 5
	CardTypeIsTrips       = 4
	CardTypeIsTwoPairs    = 3
	CardTypeIsOnePair     = 2
	CardTypeIsNormal      = 1
)

type SortedCards struct {
	Cards     []Card
	CardsType int
}

func NewSortedCards(cards []Card) *SortedCards {
	sortedCards := &SortedCards{
		Cards: cards,
	}
	// first sort only by card value asc
	sortedCards.sortByValue()
	// second calc the cards type
	sortedCards.caclMaxType()
	// finally sort by card type and value desc
	sortedCards.sortByTypeValue()

	return sortedCards
}

func (s *SortedCards) sortByValue() {
	sort.Slice(s.Cards, func(i, j int) bool {
		return s.Cards[i].GetIntValue() < s.Cards[j].GetIntValue()
	})
}

// calc the cards type
func (s *SortedCards) caclMaxType() {
	var c Comparator
	var temp [5]Card
	copy(temp[:], s.Cards)
	if c.IsStraightFlush(temp) {
		s.CardsType = CardTypeStraightFlush
		return
	}
	if c.IsFourOfAType(temp) {
		s.CardsType = CardTypeIsFour
		return
	}
	if c.IsFullHouse(temp) {
		s.CardsType = CardTypeIsFullHouse
		return
	}
	if c.IsFlush(temp) {
		s.CardsType = CardTypeIsFlush
		return
	}
	if c.IsStraight(temp) {
		s.CardsType = CardTypeIsStraight
		return
	}
	if c.IsTrips(temp) {
		s.CardsType = CardTypeIsTrips
		return
	}
	if c.IsTwoPairs(temp) {
		s.CardsType = CardTypeIsTwoPairs
		return
	}
	if c.IsOnePair(temp) {
		s.CardsType = CardTypeIsOnePair
		return
	}
	s.CardsType = CardTypeIsNormal
	return
}

// sort cards by type and value
func (s *SortedCards) sortByTypeValue() {
	switch s.CardsType {
	case CardTypeStraightFlush, CardTypeIsFlush:
		{
			// do nothing
		}
	default:
		{
			// 从大到小重新进行排序
			s.bubbleSameCards()
		}
	}
}

// bubble all the same cards
func (s *SortedCards) bubbleSameCards() {
	counter := make(map[int][]Card)
	// counter every value length
	for _, card := range s.Cards {
		counter[card.GetIntValue()] = append(counter[card.GetIntValue()], card)
	}
	// copy all the value into temp slice
	temp := make([][]Card, 0)
	for _, v := range counter {
		temp = append(temp, v)
	}
	sort.Slice(temp, func(i, j int) bool {
		// the longer the winner
		if len(temp[i]) > len(temp[j]) {
			return true
		} else if len(temp[i]) == len(temp[j]) {
			// the bigger the winner
			if temp[i][0].GetIntValue() > temp[j][0].GetIntValue() {
				return true
			} else {
				return false
			}
		}
		return false
	})
	s.Cards = nil
	// copy all the card to s.Cards
	for _, v := range temp {
		s.Cards = append(s.Cards, v...)
	}
}
