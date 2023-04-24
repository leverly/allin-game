package card

type CardValue struct {
	stringValue string
	intValue    int
}

type Card struct {
	suit  string
	value CardValue
}

func (c *Card) GetIntValue() int {
	return c.value.intValue
}

func (c *Card) GetStrValue() string {
	return c.value.stringValue
}

func (c *Card) ValueEqual(other Card) bool {
	return c.value.intValue == other.value.intValue
}

func (c *Card) TypeEqual(other Card) bool {
	return c.suit == other.suit
}
