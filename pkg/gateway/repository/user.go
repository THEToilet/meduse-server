package repository

import (
	"meduse-server/pkg/domain/model"
	"meduse-server/pkg/domain/repository"
	"sync"
)

var _ repository.UserRepository = &UserRepository{}

var (
	UserList = &sync.Map{}
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u UserRepository) Find(userID string) (*model.User, error) {
	userInfo, ok := UserList.Load(userID)
	if !ok {
		return nil, model.ErrUserNotFound
	}
	v, ok := userInfo.(model.User)
	if !ok {
		return nil, model.ErrUserNotFound
	}
	return &v, nil
}

func (u UserRepository) FindAll() []*model.User {
	var userList []*model.User
	UserList.Range(func(key, value interface{}) bool {
		v, ok := value.(model.User)
		if !ok {
			return false
		}
		userList = append(userList, &v)
		return true
	})
	return userList
}

func (u UserRepository) Save(user model.User) error {
	_, ok := UserList.Load(user.UserID)
	if ok {
		return model.ErrUserAlreadyExisted
	}
	UserList.Store(user.UserID, user)
	return nil
}

func (u UserRepository) Update(user model.User) error {
	_, ok := UserList.Load(user.UserID)
	if !ok {
		return model.ErrUserNotFound
	}
	UserList.Store(user.UserID, user)
	return nil
}

func (u UserRepository) Delete(userID string) error {
	_, ok := UserList.Load(userID)
	if !ok {
		return model.ErrUserNotFound
	}
	UserList.Delete(userID)
	return nil
}
