package model

type RoomState string

const (
	Making                 = RoomState("Making")
	Matching               = RoomState("Matching")
	Progressing            = RoomState("Progressing")
	WaitingForHost         = RoomState("WaitingForHost")
	WaitingForParticipants = RoomState("WaitingForParticipants")
)
