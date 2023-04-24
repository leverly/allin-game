package game

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
	r.lock()

	// random the start player index
	start := rand.Int() % len(r.Players)

	// start the gambling many times
	for i := 0; i < 10; i++ {
		gambling := NewGame(start%len(r.Players), r.Players)
		gambling.Start()
		start++
	}
	// the player can join/leave right now
	r.unlock()
	return true
}

// lock the room and arrange position
func (r *Room) lock() {
	r.Locked = true
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(r.Players), func(i, j int) {
		r.Players[i], r.Players[j] = r.Players[j], r.Players[i]
	})
}

// unlock the room
func (r *Room) unlock() {
	r.Locked = false
}

// enter the room when unlocked
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

// leave the room when unlocked
func (r *Room) Leave(name string) bool {
	if r.Locked {
		return false
	}
	// find the player the remove
	for i := 0; i < len(r.Players); i++ {
		if r.Players[i].Name == name {
			r.Players = append(r.Players[:i], r.Players[i+1:]...)
			return true
		}
	}
	return false
}
