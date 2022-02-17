package model

import "errors"

var (
	ErrTest = errors.New("test")
	ErrUserNotFound         = errors.New("user not found")
	ErrUserAlreadyExisted   = errors.New("user already existed")
	ErrCannotGenerateUserID = errors.New("can not generate userID")
	ErrCannotGenerateRoomID = errors.New("can not generate roomID")
)
