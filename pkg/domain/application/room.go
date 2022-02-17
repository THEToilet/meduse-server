package application

import (
	"github.com/rs/zerolog"
	"meduse-server/pkg/gateway/repository"
)

type RoomUseCase struct {
	logger *zerolog.Logger
}

func NewRoomUseCase(*repository.RoomRepository) *RoomUseCase {
	return &RoomUseCase{

	}
}
