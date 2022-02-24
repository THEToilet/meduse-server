package repository

import (
	"github.com/google/uuid"
	"meduse-server/pkg/domain/model"
	"meduse-server/pkg/domain/repository"
	"sync"
)

var _ repository.RoomRepository = &RoomRepository{}

var (
	RoomList = &sync.Map{}
)

type RoomRepository struct {
}

func NewRoomRepository() (*RoomRepository, error) {
	roomID, err := uuid.NewUUID()
	if err != nil {
		return nil, model.ErrCannotGenerateRoomID
	}
	RoomList.Store(roomID.String(), model.NewRoom(roomID.String()))
	return &RoomRepository{}, nil
}

func (r RoomRepository) Find(id string) (*model.User, error) {
	panic("implement me")
}

func (r RoomRepository) Save(user model.User) error {
	panic("implement me")
}

func (r RoomRepository) Update(user model.User) error {
	panic("implement me")
}

func (r RoomRepository) FindAll() []*model.Room {
	panic("implement me")
}

func (r RoomRepository) DeleteUser(roomID string, userID string) error {
	panic("implement me")
}
