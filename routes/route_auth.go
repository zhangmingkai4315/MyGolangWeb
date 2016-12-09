package routes

import (
	"data"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"templates"
	"utils"
)

//func Authenticate(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	user, _ := data.UserByEmail(r.PostFormValue("email"))
//	if user.Password == data.Encrypt(r.PostFormValue("password")) {
//		session := user.CreateSession()
//		cookie := http.Cookie{
//			Name:     "_cookie",
//			Value:    session.Uuid,
//			HttpOnly: true,
//		}
//		http.SetCookie(w, &cookie)
//		http.Redirect(w, r, "/", 302)
//	} else {
//		http.Redirect(w, r, "/login", 302)
//	}
//}

// Do the first check work and collect data from cookie and session.
func statusCollection(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	userSession := &data.Session{}
	authChecker, _ := data.ValidateSession(w, r, userSession)
	render_data := map[string]interface{}{
		"Auth": authChecker,
	}
	flash_message, _ := ShowFlashMessage(w, r)
	flash_message_type := strings.Split(flash_message, "||")
	if len(flash_message_type) == 0 {
		render_data["FlashSuccess"] = false
	} else {
		render_data["FlashSuccess"] = true
		flash_message = flash_message_type[0]
	}
	if len(flash_message) != 0 {
		render_data["Flash"] = flash_message
	}
	return render_data
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	render_data := statusCollection(w, r)
	if render_data["Auth"] == true {
		http.Redirect(w, r, "/logout", http.StatusTemporaryRedirect)
	} else {
		templates.RenderTemplate(w, "login", render_data)
	}
}

func PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	userSession := &data.Session{User: data.User{}}
	data.ValidateSession(w, r, userSession)
	email := r.FormValue("email")
	password := r.FormValue("password")
	// do some serve validation!
	if len(email) == 0 || len(password) == 0 {
		SetFlashMessage(w, r, []byte("用户名或密码输入不能为空，请重新输入"))
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	err := data.Db.QueryRow("select id,user_name,user_password from users where user_email=? limit 1",
		email).Scan(&userSession.User.Id, &userSession.User.Username, &userSession.User.Password)
	switch {
	case err == sql.ErrNoRows:
		SetFlashMessage(w, r, []byte("输入邮箱地址不存在，请重新输入"))
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	case err != nil:
		SetFlashMessage(w, r, []byte("服务器暂时无法处理请求，请稍后再试"))
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	default:
		if bcrypt.CompareHashAndPassword([]byte(userSession.User.Password), []byte(password)) != nil {
			SetFlashMessage(w, r, []byte("输入邮箱地址和密码不匹配"))
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		data.UpdateSession(userSession.Id, userSession.User.Id)
		SetFlashMessage(w, r, []byte("您已成功登入系统||ok"))
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// validate session first for render different page view.

	render_data := statusCollection(w, r)
	// if user has already login, we will redirect the user to logout status.
	if render_data["Auth"] == true {
		http.Redirect(w, r, "/logout", http.StatusTemporaryRedirect)
	} else {
		templates.RenderTemplate(w, "register", render_data)
	}
}

func PostRegisterHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Redirect(w, r, "/register", http.StatusInternalServerError)
		return
	}
	username := r.FormValue("username")
	email := r.FormValue("email")

	//maybe we can get user back to the original url. Try it later/
	//pageGuid := r.FormValue("referer")
	password := r.FormValue("password")
	password2 := r.FormValue("password2")
	// do more validation
	// ...
	// ...
	if password != password2 {
		SetFlashMessage(w, r, []byte("两次密码输入不一致"))
		http.Redirect(w, r, "/register", http.StatusFound)
		return
	}

	// when username or email has already occupied.
	rows, err := data.Db.Query("Select user_name,user_email from users where user_name=? or user_email=? limit 1", username, email)
	if rows.Next() {
		SetFlashMessage(w, r, []byte("用户名或邮箱已被占用，请重新输入注册信息"))
		http.Redirect(w, r, "/register", http.StatusFound)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.ErrorLog.Println(err.Error())
		SetFlashMessage(w, r, []byte("服务器暂时无法处理请求，请稍后再试"))
		http.Redirect(w, r, "/register", http.StatusInternalServerError)
		return
	}
	_, err = data.Db.Exec("Insert into users set user_name=?, user_email=?, user_password=?", username, email, string(hashedPassword))
	if err != nil {
		utils.ErrorLog.Println(err.Error())
		SetFlashMessage(w, r, []byte("服务器暂时无法处理请求，请稍后再试"))
		http.Redirect(w, r, "/register", http.StatusInternalServerError)
		return
	} else {
		SetFlashMessage(w, r, []byte("您已注册成功，请先使用账号密码登陆||ok"))
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// validate session first for render different page view.
	data.DeleteSession(w, r)
	SetFlashMessage(w, r, []byte("您已退出登入状态||OK"))
	http.Redirect(w, r, "/", http.StatusFound)
}
