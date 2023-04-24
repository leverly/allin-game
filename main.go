package main

import "allin/game"

func main() {
	gameMall := game.NewGameMall()
	room := gameMall.CreateRoom("room", "password")

	player1 := game.Player{Name: "player1", RemainCoin: game.CHIPS, LoanCount: 1, Active: true}
	room.Enter("password", player1)
	player2 := game.Player{Name: "player2", RemainCoin: game.CHIPS, LoanCount: 1, Active: true}
	room.Enter("password", player2)
	player3 := game.Player{Name: "player3", RemainCoin: game.CHIPS, LoanCount: 1, Active: true}
	room.Enter("password", player3)
	player4 := game.Player{Name: "player4", RemainCoin: game.CHIPS, LoanCount: 1, Active: true}
	room.Enter("password", player4)
	player5 := game.Player{Name: "player5", RemainCoin: game.CHIPS, LoanCount: 1, Active: true}
	room.Enter("password", player5)

	room.Start("password")
}
