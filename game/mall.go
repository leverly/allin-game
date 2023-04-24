package game

type Mall struct {
	Rooms map[string]*Room
}

func NewGameMall() *Mall {
	return &Mall{
		Rooms: make(map[string]*Room),
	}
}

func (m *Mall) CreateRoom(name, password string) *Room {
	_, find := m.Rooms[name]
	if find {
		return nil
	}
	room := &Room{Name: name, Password: password}
	m.Rooms[name] = room
	return room
}

func (m *Mall) ListRoomName() []string {
	var list []string
	for _, room := range m.Rooms {
		list = append(list, room.Name)
	}
	return list
}
