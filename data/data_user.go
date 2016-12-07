package data

import (
	"strconv"
	"time"
)

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	Status    string
	Phone     string
}

func (user *User) CreateSession() (session Session) {
	session = Session{
		Id:        1,
		Uuid:      "uuid1",
		Email:     user.Email,
		UserId:    strconv.Itoa(user.Id),
		CreatedAt: time.Now(),
	}
	return
}

func Encrypt(password string) string {
	return password
}

func UserByEmail(email string) (user User, err error) {
	user = User{
		Id:        1,
		Name:      "TestUser",
		Email:     email,
		Password:  email,
		CreatedAt: time.Now(),
	}
	return
}
