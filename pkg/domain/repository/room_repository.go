package repository

import "meduse-server/pkg/domain/model"

type RoomRepository interface {
	Find(id string) (*model.User, error)
	Save(user model.User) error
	Update(user model.User) error

	FindAll() []*model.Room
	DeleteUser(roomID string, userID string) error
}
