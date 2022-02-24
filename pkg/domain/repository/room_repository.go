package repository

import "meduse-server/pkg/domain/model"

type RoomRepository interface {
	Save(user model.User) error
	Update(user model.User) error
	Find(id string) (*model.User, error)
	FindAll() []*model.User
	DeleteUser(userID string) error
	DeleteAll()
}
