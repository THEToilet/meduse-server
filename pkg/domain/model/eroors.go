package model

import "errors"

var (
	ErrTest                         = errors.New("test")
	ErrUserNotFound                 = errors.New("user not found")
	ErrUserAlreadyExisted           = errors.New("user already existed")
	ErrUserConnectionNotFound       = errors.New("user connection not found")
	ErrUserConnectionAlreadyExisted = errors.New("user connection already existed")
	ErrCannotGenerateUserID         = errors.New("can not generate userID")
	ErrRoomIsFull                   = errors.New("the room is full")
)
