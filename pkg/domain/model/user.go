package model

type User struct {
	UserID   string
	UserName string
}

func NewUser(userID string, userName string) *User {
	return &User{
		UserID:   userID,
		UserName: userName,
	}
}
