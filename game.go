package main

import (
	"fmt"
	"math/rand"
)

var CHIPS int = 1000

type Game struct {
	cards      Cards
	roomPlayer []Player
	startIndex int
	deskCards  []Card
}

func NewGame(start int, player []Player) *Game {
	// reset last gambling info except the remain coin
	for i := 0; i < len(player); i++ {
		if player[i].RemainCoin == 0 {
			// init the coin
			player[i].Loan(CHIPS)
		}
		player[i].Clean()
	}

	// new game starter
	game := Game{roomPlayer: player, startIndex: start}
	game.cards.Init()
	return &game
}

func (g *Game) gambling(round string) {
	count := len(g.roomPlayer)
	maxCoin := 0
	start := g.startIndex
	maxIndex := start
	for i := 0; i < count; i++ {
		// finish the gambling
		if g.getActiveCount() < 2 {
			return
		}
		coin := g.roomPlayer[start%count].Action(round, maxCoin)
		if coin > maxCoin {
			maxCoin = coin
			maxIndex = start % count
		}
		start++
	}
	// before maxIndex player reaction
	for i := g.startIndex; i%count != maxIndex; i++ {
		// finish the gambling
		if g.getActiveCount() < 2 {
			return
		}
		g.roomPlayer[start%count].Reaction(round, maxCoin)
		start++
	}
}

func (g *Game) kickoff() {
	// 洗牌
	g.cards.Shuffle()
	// 发2张手牌
	count := len(g.roomPlayer)
	for i := 0; i < 2; i++ {
		start := g.startIndex
		for j := 0; j < count; j++ {
			g.roomPlayer[start%count].AddCard(g.cards.Next())
			start++
		}
	}
	// 发5张底牌
	g.cards.Next()
	for i := 0; i < 3; i++ {
		g.deskCards = append(g.deskCards, g.cards.Next())
	}
	g.cards.Next()
	g.deskCards = append(g.deskCards, g.cards.Next())
	g.cards.Next()
	g.deskCards = append(g.deskCards, g.cards.Next())
}

func (g *Game) Start() {
	// 准备开始
	g.kickoff()

	// 开始押注
	g.gambling("first round")

	count := len(g.roomPlayer)
	// 翻三张底牌
	for i := 0; i < 3; i++ {
		card := g.deskCards[i]
		for j := 0; j < count; j++ {
			g.roomPlayer[j].AddCard(card)
		}
	}
	g.gambling("second round")

	// 翻一张底牌
	card := g.deskCards[3]
	for j := 0; j < count; j++ {
		g.roomPlayer[j].AddCard(card)
	}
	g.gambling("third round")

	// 翻最后一张底牌
	card = g.deskCards[4]
	for j := 0; j < count; j++ {
		g.roomPlayer[j].AddCard(card)
	}
	g.gambling("last round")

	// 最后的清算
	g.close()
}

func (g *Game) close() {
	count := len(g.roomPlayer)
	winner := g.getWinner()
	for i := 0; i < count; i++ {
		g.roomPlayer[winner].Win(g.roomPlayer[i].curCoins)
		if i != winner {
			fmt.Println("loser:", g.roomPlayer[i].String())
		}
	}
	fmt.Println("winner:", g.roomPlayer[winner].String())
}

func (g *Game) getActiveCount() int {
	var count int
	for _, player := range g.roomPlayer {
		if player.Active {
			count++
		}
	}
	return count
}

// TODO compare the cards
func (g *Game) getWinner() int {
	counter := rand.Int()%len(g.roomPlayer) + 1
	count := 0
	for {
		for index, player := range g.roomPlayer {
			if player.Active {
				count++
				if count%counter == 0 {
					return index
				}
			}
		}
	}
}
