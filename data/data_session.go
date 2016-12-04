package data

import (
	"time"
	"net/http"
	"errors"
)

type Session struct {
	Id int
	Uuid string
	Email string
	UserId string
	CreatedAt time.Time
}

func (s *Session)Check()(ok bool,err error){
	ok=true
	return
}
func GetSession (w http.ResponseWriter,r *http.Request)(session Session,err error){
	cookie,err:=r.Cookie("_cookie")
	if err==nil{
		session=Session{Uuid:cookie.Value}
		if ok,_:=session.Check();!ok{
			err = errors.New("Invalid session")
		}
	}
	return
}