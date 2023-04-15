package main

func main() {
	gameMall := NewGameMall()
	room := gameMall.CreateRoom("room", "password")

	player1 := Player{Name: "player1", RemainCoin: CHIPS, LoanCount: 1, Active: true}
	room.Enter("password", player1)
	player2 := Player{Name: "player2", RemainCoin: CHIPS, LoanCount: 1, Active: true}
	room.Enter("password", player2)
	player3 := Player{Name: "player3", RemainCoin: CHIPS, LoanCount: 1, Active: true}
	room.Enter("password", player3)
	player4 := Player{Name: "player4", RemainCoin: CHIPS, LoanCount: 1, Active: true}
	room.Enter("password", player4)
	player5 := Player{Name: "player5", RemainCoin: CHIPS, LoanCount: 1, Active: true}
	room.Enter("password", player5)

	room.Start("password")
}
