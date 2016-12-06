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
}
func main() {
	r := mux.NewRouter()

	//mux := http.NewServeMux()
	//files := http.FileServer(http.Dir("./public"))
	//mux.Handle("/static/", http.StripPrefix("/static", files))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	//handle '/' and any unknown routes
	r.HandleFunc("/", routes.IndexHandler)
	r.HandleFunc("/error", routes.ErrHandler)

	listenAddr := viper.GetString(utils.Env + ".webserver.listenAt")
	server := &http.Server{
		Addr:    listenAddr,
		Handler: r,
	}
	utils.InfoLog.Println("Start WebServer at", listenAddr)
	if err := server.ListenAndServe(); err != nil {
		utils.ErrorLog.Println("Server start error:", err)
	}
}
