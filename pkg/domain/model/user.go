package model

type User struct {
	UserID         string
	UserName       string
	ControllerName ControllerName
	IsHost         bool
}

func NewUser(userID string, userName string, isHost bool) *User {
	return &User{
		UserID:         userID,
		UserName:       userName,
		ControllerName: CON0,
		IsHost:         isHost,
	}
}
