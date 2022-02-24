package model

type Room struct {
	RoomID      string
	RoomState   RoomState
	Host        User
	Participant [8]User
}

func NewRoom(roomID string) *Room {
	return &Room{
		RoomID:      roomID,
		Host:        User{},
		RoomState:   Matching,
		Participant: [8]User{},
	}
}
