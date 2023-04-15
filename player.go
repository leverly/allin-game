package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	Name       string
	RemainCoin int
	LoanCount  int
	Active     bool

	curCoins []int
	curCards []Card
}

func (p *Player) Loan(coin int) {
	p.LoanCount += 1
	p.RemainCoin += coin
}

// clean last round data
func (p *Player) Clean() {
	p.curCoins = nil
	p.curCards = nil
	p.Active = true
}

func (p *Player) String() string {
	return fmt.Sprintf("name[%s], vote[%d], active[%v], loan[%d], remain[%d], win[%d]",
		p.Name, p.getVote(), p.Active, p.LoanCount, p.RemainCoin, p.RemainCoin-p.LoanCount*1000)
}

func (p *Player) Action(round string, max int) int {
	// already quit
	if p.Active == false {
		return 0
	}
	// random quit
	if rand.Int()%10 == 0 {
		p.Active = false
		// fmt.Println(round, "act:", p.String())
		return 0
	}
	// follow up
	coin := rand.Int() % 100
	if coin < max {
		coin = max
	} else {
		max = coin
	}
	// all in
	if p.RemainCoin < coin {
		coin = p.RemainCoin
	}
	// record the coin chain
	p.curCoins = append(p.curCoins, coin)
	p.RemainCoin = p.RemainCoin - coin
	// fmt.Println(round, "act:", p.String())
	return coin
}

func (p *Player) Reaction(round string, max int) int {
	// already quit
	if p.Active == false {
		return 0
	}
	// random quit
	if rand.Int()%10 == 0 {
		p.Active = false
		// fmt.Println(round, "react:", p.String())
		return 0
	}
	// follow the max
	coin := max - p.curCoins[len(p.curCoins)-1]
	// all in
	if p.RemainCoin < coin {
		coin = p.RemainCoin
	}
	// modify the coin chain
	p.curCoins[len(p.curCoins)-1] += coin
	p.RemainCoin = p.RemainCoin - coin
	// fmt.Println(round, "react:", p.String())
	return coin
}

func (p *Player) AddCard(card Card) {
	p.curCards = append(p.curCards, card)
}

func (p *Player) Win(coins []int) {
	for _, coin := range coins {
		p.RemainCoin += coin
	}
}

func (p *Player) getVote() int {
	var amount int
	for _, coin := range p.curCoins {
		amount += coin
	}
	return amount
}
