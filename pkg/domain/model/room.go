package model

type Room struct {
	RoomID      string
	RoomState   RoomState
	Host        User
	Participant [3]User
}

func NewRoom(roomID string) *Room {
	return &Room{
		RoomID:      roomID,
		Host:        User{},
		RoomState:   Matching,
		Participant: [3]User{},
	}
}
