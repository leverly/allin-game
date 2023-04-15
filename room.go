package main

import (
	"math/rand"
	"time"
)

type Room struct {
	Name     string
	Password string
	Players  []Player
	Locked   bool
}

func (r *Room) Start(password string) bool {
	if r.Password != password {
		return false
	}
	if len(r.Players) <= 1 {
		return false
	}
	// lock room and shuffle the player position
	r.lockRoom()
	// random the start player index
	start := rand.Int() % len(r.Players)

	for i := 0; i < 5; i++ {
		gambling := NewGame(start%len(r.Players), r.Players)
		gambling.Start()
		start++
	}
	return true
}

// lock the room and arrange position
func (r *Room) lockRoom() {
	r.Locked = true
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(r.Players), func(i, j int) {
		r.Players[i], r.Players[j] = r.Players[j], r.Players[i]
	})
}

func (r *Room) Enter(password string, player Player) bool {
	if r.Locked {
		return false
	}
	if r.Password == password {
		r.Players = append(r.Players, player)
		return true
	}
	return false
}
