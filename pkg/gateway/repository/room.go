package repository

import (
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

func NewRoomRepository() *RoomRepository {
	return &RoomRepository{}
}

func (r RoomRepository) Save(user model.User) error {
	_, ok := RoomList.Load(user.UserID)
	if ok {
		return model.ErrUserAlreadyExisted
	}
	RoomList.Store(user.UserID, user)
	return nil
}

func (r RoomRepository) Update(user model.User) error {
	_, ok := RoomList.Load(user.UserID)
	if !ok {
		return model.ErrUserNotFound
	}
	RoomList.Store(user.UserID, user)
	return nil
}

func (r RoomRepository) Find(userID string) (*model.User, error) {
	userInfo, ok := RoomList.Load(userID)
	if !ok {
		return nil, model.ErrUserNotFound
	}
	v, ok := userInfo.(model.User)
	if !ok {
		return nil, model.ErrUserNotFound
	}
	return &v, nil
}

func (r RoomRepository) FindAll() []*model.User {
	var userList []*model.User
	RoomList.Range(func(key, value interface{}) bool {
		v, ok := value.(model.User)
		if !ok {
			return false
		}
		userList = append(userList, &v)
		return true
	})
	return userList
}

func (r RoomRepository) DeleteUser(userID string) error {
	_, ok := RoomList.Load(userID)
	if !ok {
		return model.ErrUserNotFound
	}
	RoomList.Delete(userID)
	return nil
}

func (r RoomRepository) DeleteAll() {
	RoomList.Range(func(key, value interface{}) bool {
		RoomList.Delete(key)
		return true
	})
}
