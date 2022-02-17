package application

import (
	"github.com/rs/zerolog"
	"meduse-server/pkg/gateway/repository"
)

type UserUseCase struct {
	logger *zerolog.Logger
}

func NewUserUseCase(*repository.UserRepository) *UserUseCase {
	return &UserUseCase{

	}
}
