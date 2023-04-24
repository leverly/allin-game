package card

// input cards already sorted by card's int value
type Comparator struct {
}

// compare two cards by type if the same then compare with value
func (c *Comparator) IsGreater(A, B *SortedCards) int {
	if A.CardsType > B.CardsType {
		return 1
	} else if A.CardsType < B.CardsType {
		return -1
	}
	// compare the values the bigger win
	for i := len(A.Cards) - 1; i >= 0; i-- {
		if A.Cards[i].GetIntValue() > B.Cards[i].GetIntValue() {
			return 1
		} else if A.Cards[i].GetIntValue() < B.Cards[i].GetIntValue() {
			return -1
		}
	}
	return 0
}

// 同花顺
func (c *Comparator) IsStraightFlush(cards [5]Card) bool {
	return c.IsFlush(cards) && c.IsStraight(cards)
}

// 四个
func (c *Comparator) IsFourOfAType(cards [5]Card) bool {
	return cards[1].ValueEqual(cards[2]) && cards[2].ValueEqual(cards[3]) &&
		(cards[3].ValueEqual(cards[4]) || cards[0].ValueEqual(cards[1]))
}

// 葫芦
func (c *Comparator) IsFullHouse(cards [5]Card) bool {
	return cards[0].ValueEqual(cards[1]) && cards[3].ValueEqual(cards[4]) &&
		(cards[2].ValueEqual(cards[0]) || cards[2].ValueEqual(cards[4]))
}

// 同花
func (c *Comparator) IsFlush(cards [5]Card) bool {
	return cards[0].TypeEqual(cards[1]) && cards[1].TypeEqual(cards[2]) &&
		cards[2].TypeEqual(cards[3]) && cards[3].TypeEqual(cards[4])
}

// 顺子
func (c *Comparator) IsStraight(cards [5]Card) bool {
	if cards[0].GetIntValue() == cards[1].GetIntValue()-1 &&
		cards[1].GetIntValue() == cards[2].GetIntValue()-1 &&
		cards[2].GetIntValue() == cards[3].GetIntValue()-1 &&
		(cards[3].GetIntValue() == cards[4].GetIntValue()-1 ||
			(cards[3].GetStrValue() == "5" && cards[4].GetStrValue() == "A")) {
		return true
	}
	return false
}

// 三个
func (c *Comparator) IsTrips(cards [5]Card) bool {
	return cards[0].ValueEqual(cards[2]) || cards[1].ValueEqual(cards[3]) || cards[2].ValueEqual(cards[4])
}

// 俩对
func (c *Comparator) IsTwoPairs(cards [5]Card) bool {
	return (cards[0].ValueEqual(cards[1]) && cards[2].ValueEqual(cards[3])) ||
		(cards[0].ValueEqual(cards[1]) && cards[3].ValueEqual(cards[4])) ||
		(cards[1].ValueEqual(cards[2]) && cards[3].ValueEqual(cards[4]))
}

// 一对
func (c *Comparator) IsOnePair(cards [5]Card) bool {
	return cards[0].ValueEqual(cards[1]) || cards[1].ValueEqual(cards[2]) ||
		cards[2].ValueEqual(cards[3]) || cards[3].ValueEqual(cards[4])
}
