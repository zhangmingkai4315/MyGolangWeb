package routes

import (
	"encoding/base64"
	"net/http"
	"time"
)

func encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}
func SetFlashMessage(w http.ResponseWriter, r *http.Request, msg []byte) {
	c := http.Cookie{
		Name:  "flash_message",
		Value: encode(msg),
	}
	http.SetCookie(w, &c)
}

func ShowFlashMessage(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	c, err := r.Cookie("flash_message")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, nil
		} else {
			return nil, err

		}
	}
	value, err := decode(c.Value)
	if err != nil {
		return nil, err
	}
	dc := &http.Cookie{Name: "flash_message", MaxAge: -1, Expires: time.Unix(1, 0)}
	http.SetCookie(w, dc)
	return value, nil
}
