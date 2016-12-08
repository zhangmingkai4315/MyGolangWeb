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

func ShowFlashMessage(w http.ResponseWriter, r *http.Request) (value string, err error) {
	c, err := r.Cookie("flash_message")
	if err != nil {
		if err == http.ErrNoCookie {
			return
		} else {
			return

		}
	}
	value_byte, err := decode(c.Value)
	if err != nil {
		return
	}
	value = string(value_byte[:])
	dc := &http.Cookie{Name: "flash_message", MaxAge: -1, Expires: time.Unix(1, 0)}
	http.SetCookie(w, dc)
	return
}
