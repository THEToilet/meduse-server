package model

type User struct {
	UserID   string
	UserName string
	IsHost   bool
}

func NewUser(userID string, userName string, isHost bool) *User {
	return &User{
		UserID:   userID,
		UserName: userName,
		IsHost:   isHost,
	}
}
