package routes

import (
	"data"
	"net/http"
	"templates"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, _ := data.UserByEmail(r.PostFormValue("email"))
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session := user.CreateSession()
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	_, sessionErr := data.GetSession(w, r)
	render_data := map[string]interface{}{}
	flash_message, _ := ShowFlashMessage(w, r)
	if sessionErr == nil {
		http.Redirect(w, r, "/logout", http.StatusTemporaryRedirect)
	} else {
		render_data["Auth"] = false
		if len(flash_message) != 0 {
			render_data["Flash"] = flash_message
		}
		templates.RenderTemplate(w, "login", render_data)
	}
}

func PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	_, sessionErr := data.GetSession(w, r)
	//render_data := map[string]interface{}{"Auth": false}
	if sessionErr == nil {
		// user session exist, maybe user has been authenticated before. so logout first!
		http.Redirect(w, r, "/logout", http.StatusTemporaryRedirect)
		return
	}
	// user session not exist. so try to get user post data
	err := r.ParseForm()
	if err != nil {
		SetFlashMessage(w, r, []byte("服务器暂时无法处理请求，请稍后再试"))
		http.Redirect(w, r, "/login", http.StatusInternalServerError)
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	// do some serve validation!
	if len(email) == 0 || len(password) == 0 {
		SetFlashMessage(w, r, []byte("用户名或密码输入不能为空，请重新输入"))
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	// compare user post data with user database information.
	if email == password {
		//save user session and success login with redirect to home page.(later try to redirect to the url before login)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	} else {
		SetFlashMessage(w, r, []byte("用户名或密码输入不正确，请重新输入"))
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	_, sessionErr := data.GetSession(w, r)
	render_data := map[string]interface{}{}
	flash_message, _ := ShowFlashMessage(w, r)
	// if user has already login, we will redirect the user to logout status.
	if sessionErr == nil {
		http.Redirect(w, r, "/logout", http.StatusTemporaryRedirect)
	} else {
		render_data["Auth"] = false
		if len(flash_message) != 0 {
			render_data["Flash"] = flash_message
		}
		templates.RenderTemplate(w, "register", render_data)
	}
}

func PostRegisterHandler(w http.ResponseWriter, r *http.Request) {

}
