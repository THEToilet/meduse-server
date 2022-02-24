package application

import (
	"github.com/rs/zerolog"
	"meduse-server/pkg/domain/model"
	"meduse-server/pkg/domain/repository"
)

type RoomUseCase struct {
	RoomState      model.RoomState
	Host           model.User
	Controller     []model.User
	roomRepository repository.RoomRepository
	logger         *zerolog.Logger
}

func NewRoomUseCase(roomRepository repository.RoomRepository, logger *zerolog.Logger) *RoomUseCase {
	return &RoomUseCase{
		roomRepository: roomRepository,
		RoomState:      model.Making,
		logger:         logger,
	}
}

const (
	MaxController = 4
	MaxUser       = 8
)

func (r *RoomUseCase) RegisterUser() error {
	// TODO:部屋が8人以上だと拒否する
	if len(r.roomRepository.FindAll()) > MaxUser {
		return model.ErrRoomIsFull
	}
	return nil
}

// UpdateUser Update 名前の更新
func (r *RoomUseCase) UpdateUser(user model.User) error {
	err := r.roomRepository.Update(user)
	return err
}

func (r *RoomUseCase) DeleteUser(userID string) error {
	err := r.roomRepository.DeleteUser(userID)
	return err
}

func (r *RoomUseCase) DeleteAllUser() {
	r.roomRepository.DeleteAll()
}

func (r *RoomUseCase) ReturnToSpectators(userID string) {
	for i, user := range r.Controller {
		if user.UserID == userID {
			r.Controller[i] = model.User{}
		}
	}
}

func (r *RoomUseCase) SpectatorMoveToController(userID string, controllerNumber int) error {
	user, err := r.roomRepository.Find(userID)
	if err != nil {
		return model.ErrUserNotFound
	}
	r.Controller[controllerNumber-1] = *user
	return nil
}

func (r *RoomUseCase) GetAllControllerUser() []model.User {
	return r.Controller
}
