package application

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"meduse-server/pkg/domain/model"
	"meduse-server/pkg/domain/repository"
)

type RoomUseCase struct {
	roomState      model.RoomState
	host           model.User
	controller     []model.User
	roomRepository repository.RoomRepository
	logger         *zerolog.Logger
}

func NewRoomUseCase(roomRepository repository.RoomRepository, logger *zerolog.Logger) *RoomUseCase {
	return &RoomUseCase{
		roomRepository: roomRepository,
		roomState:      model.Making,
		logger:         logger,
	}
}

const (
	MaxController = 4
	MaxUser       = 8
)

func (r *RoomUseCase) RegisterUser() (string, error) {
	// TODO:部屋が8人以上だと拒否する
	if len(r.roomRepository.FindAll()) > MaxUser {
		return "", model.ErrRoomIsFull
	}
	userID, err := uuid.NewUUID()
	if err != nil {
		return "", model.ErrCannotGenerateUserID
	}
	return userID.String(), nil
}

// UpdateUser Update 名前の更新
func (r *RoomUseCase) UpdateUser(user model.User) error {
	err := r.roomRepository.Update(user)
	return err
}

func (r *RoomUseCase) GetHostUser() (model.User, error) {
	return r.host, nil
}

func (r *RoomUseCase) DeleteUser(userID string) error {
	err := r.roomRepository.DeleteUser(userID)
	return err
}

func (r *RoomUseCase) DeleteAllUser() {
	r.roomRepository.DeleteAll()
}

func (r *RoomUseCase) ReturnToSpectators(userID string) error {
	for i, user := range r.controller {
		if user.UserID == userID {
			r.controller[i] = model.User{}
		}
	}
	return nil
}

func (r *RoomUseCase) SpectatorMoveToController(userID string, controllerNumber uint) error {
	user, err := r.roomRepository.Find(userID)
	if err != nil {
		return model.ErrUserNotFound
	}
	r.controller[controllerNumber-1] = *user
	user.ControllerName = model.GetControllerName(controllerNumber)
	return nil
}

func (r *RoomUseCase) GetAllControllerUser() []model.User {
	return r.controller
}

func (r *RoomUseCase) GetAllUser() []*model.User {
	return r.roomRepository.FindAll()
}
