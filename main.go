package main

import (
	"data"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"routes"
	_ "routes"
	"templates"
	"utils"
)

func init() {
	utils.InitConfigFile()
	utils.InitLogger()
	templates.InitTemplates()
	data.InitDataBaseConnect()
	data.InitSessionStore()
}
func main() {

	router := mux.NewRouter()

	//mux := http.NewServeMux()
	//files := http.FileServer(http.Dir("./public"))
	//mux.Handle("/static/", http.StripPrefix("/static", files))

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	//handle '/' for home page
	router.HandleFunc("/", routes.IndexHandler)
	router.HandleFunc("/login", routes.LoginHandler).Methods("GET")
	router.HandleFunc("/login", routes.PostLoginHandler).Methods("POST")

	router.HandleFunc("/register", routes.RegisterHandler).Methods("GET")
	router.HandleFunc("/register", routes.PostRegisterHandler).Methods("POST")

	router.HandleFunc("/logout", routes.LogoutHandler).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(routes.NotFoundHandler)

	listenAddr := viper.GetString(utils.Env + ".webserver.listenAt")
	utils.InfoLog.Println("Start WebServer at", listenAddr)

	csrf.Secure(false)
	authkey := []byte("IDONTWANTTELLYOUTHATANDIREALLYDONTWANTYOUKNOWIT")
	println(len(authkey))
	log.Panic(http.ListenAndServe(listenAddr, csrf.Protect(authkey[0:32], csrf.Secure(false))(router)))
}
