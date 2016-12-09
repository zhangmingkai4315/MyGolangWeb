package data

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gorilla/sessions"
	"io"
	"net/http"
)

type Session struct {
	Id              string
	Authenticated   bool
	Unauthenticated bool
	User            User
}

// please replace it with some secret string
var SessionStore *sessions.CookieStore

func InitSessionStore() {
	SessionStore = sessions.NewCookieStore([]byte("golangchina"))
}

func GetSessionUID(sid string) int {
	var userid int
	err := Db.QueryRow("select user_id from sessions WHERE session_id=?", sid).Scan(&userid)
	if err != nil {
		//utils.ErrorLog.Println(err.Error())
		return 0
	}
	return userid
}

func UpdateSession(sid string, uid int) error {
	_, err := Db.Exec("insert into sessions set session_id=?, user_id=? on duplicate key update user_id=?", sid, uid, uid)
	if err != nil {
		return err
	}
	return nil
}

func GenerateSessionId() (sid string, err error) {
	sid_byte := make([]byte, 24)
	_, err = io.ReadFull(rand.Reader, sid_byte)
	if err != nil {
		return
	}
	sid = base64.URLEncoding.EncodeToString(sid_byte)
	return sid, nil
}

func ValidateSession(w http.ResponseWriter, r *http.Request, userSession *Session) (bool, string) {
	session, _ := SessionStore.Get(r, "app-session")
	if sid, valid := session.Values["sid"]; valid {
		currentUid := GetSessionUID(sid.(string))
		userSession.Id = sid.(string)
		if currentUid != 0 {
			return true, sid.(string)
		} else {
			return false, sid.(string)
		}
	} else {
		newSid, _ := GenerateSessionId()
		session.Values["sid"] = newSid
		session.Save(r, w)
		userSession.Id = newSid
		UpdateSession(newSid, 0)
		return false, newSid
	}
}

func DeleteSession(w http.ResponseWriter, r *http.Request) error {
	session, _ := SessionStore.Get(r, "app-session")
	if sid, valid := session.Values["sid"]; valid {
		_, err := Db.Exec("delete from sessions where session_id=?", sid)
		if err!=nil{
			return err
		}
		cookie := &http.Cookie{
			Name:   "app-session",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
		http.SetCookie(w, cookie)
	}
	return nil

}
