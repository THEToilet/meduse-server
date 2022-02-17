package model

type RoomState string

const (
	Matching               = RoomState("Matching")
	Progressing            = RoomState("Progressing")
	WaitingForHost         = RoomState("WaitingForHost")
	WaitingForParticipants = RoomState("WaitingForParticipants")
)
