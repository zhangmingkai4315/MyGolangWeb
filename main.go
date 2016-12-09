package main

import (
	"data"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
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
	server := &http.Server{
		Addr:    listenAddr,
		Handler: router,
	}
	utils.InfoLog.Println("Start WebServer at", listenAddr)
	if err := server.ListenAndServe(); err != nil {
		utils.ErrorLog.Println("Server start error:", err)
	}
}
