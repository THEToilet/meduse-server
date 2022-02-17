package repository

import "meduse-server/pkg/domain/model"

type UserRepository interface {
	Find(id string) (*model.User, error)
	FindAll() []*model.User
	Save(user model.User) error
	Update(user model.User) error
	Delete(userID string) error
}
