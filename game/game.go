package game

import (
	"allin/card"
	"fmt"
)

var CHIPS int = 1000

type Game struct {
	fullCards  card.FullCards
	roomPlayer []Player
	startIndex int
	deskCards  []card.Card
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
	game.fullCards.Init()
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
	g.fullCards.Shuffle()
	// 发2张手牌
	count := len(g.roomPlayer)
	for i := 0; i < 2; i++ {
		start := g.startIndex
		for j := 0; j < count; j++ {
			g.roomPlayer[start%count].AddCard(g.fullCards.Next())
			start++
		}
	}
	// 发5张底牌
	g.fullCards.Next()
	for i := 0; i < 3; i++ {
		g.deskCards = append(g.deskCards, g.fullCards.Next())
	}
	g.fullCards.Next()
	g.deskCards = append(g.deskCards, g.fullCards.Next())
	g.fullCards.Next()
	g.deskCards = append(g.deskCards, g.fullCards.Next())
}

func (g *Game) debug() {
	fmt.Println("desk cards:", card.PrintCards(g.deskCards))
	for _, p := range g.roomPlayer {
		if p.Active {
			fmt.Println(fmt.Sprintf("name:%s, type:%d, max:%s",
				p.Name, p.MaxCard.CardsType, card.PrintCards(p.MaxCard.Cards)))
		}
	}
	fmt.Println("------------------------FINISHED--------------------------------")
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

	// 结束进行结算
	g.close()
}

func (g *Game) close() {
	count := len(g.roomPlayer)
	winner := g.getWinner()
	for i := 0; i < count; i++ {
		g.roomPlayer[winner].WinCoins(g.roomPlayer[i].curCoins)
		if i != winner {
			fmt.Println("loser:", g.roomPlayer[i].String())
		}
	}
	fmt.Println("winner:", g.roomPlayer[winner].String())
	// debug check winner info
	g.debug()
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

func (g *Game) getWinner() int {
	var comparator card.Comparator
	var winner int
	var maxCard *card.SortedCards
	for index, player := range g.roomPlayer {
		if player.Active {
			// max combination of 5 cards of the player
			playerMaxCard := player.CalcMaxCards()
			// update the player info
			g.roomPlayer[index] = player
			if maxCard == nil {
				maxCard = playerMaxCard
				winner = index
				continue
			}
			// find max combination cards of all active players
			if comparator.IsGreater(playerMaxCard, maxCard) > 0 {
				maxCard = playerMaxCard
				winner = index
			}
		}
	}
	// TODO 如果多人打平，需要平分桌上其他人的筹码
	return winner
}
