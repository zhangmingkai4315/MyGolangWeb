package data

import (
	"time"
)

type User struct {
	Id       int
	Username string
	Email    string
	Password string
	UserSalt string
	JoinedAt time.Time
}
