package model

type Room struct {
	RoomID      string
	Host        User
	Participant [3]User
}

func NewRoom(roomID string) *Room {
	return &Room{
		RoomID:      roomID,
		Host:        User{},
		Participant: [3]User{},
	}
}
