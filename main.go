package main

import (
	"./data"
	"./templates"
	"fmt"
	_ "github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"routes"
)

// index function will handle the '/' query, return the threads data for users
func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		//check if the user is a anthenticted user.
		_, err := data.GetSession(w, r)
		if err == nil {
			templates.Private_templates.ExecuteTemplate(w, "layout", threads)
		} else {
			templates.Public_templates.ExecuteTemplate(w, "layout", threads)
		}
	} else {
		log.Println(err)
	}
}

//errHandler will handler the error page, maybe for 404 or something wrong.

func errHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Someone's request is broken!", r.URL.Path)
	fmt.Fprintln(w, "This is a error page!")
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static", files))

	//handle '/' and any unknown routes
	mux.HandleFunc("/", index)

	mux.HandleFunc("/error", errHandler)

	////user authentication handler
	//mux.HandleFunc("/login",login)
	//mux.HandleFunc("/logout",logout)
	//mux.HandleFunc("/signup",signup)
	//mux.HandleFunc("/signup_account",signupAccount)
	//mux.HandleFunc("/authenticate",authenticate)
	//
	////threads handler
	//mux.HandleFunc("/thread/new",newThread)
	//mux.HandleFunc("/thread/create",createThread)
	//mux.HandleFunc("/thread/post",postThread)
	//mux.HandleFunc("/thread/read",readThread)
	listenAddr := "localhost:8080"
	server := &http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}
	log.Println("Start WebServer at", listenAddr)
	log.Panic(server.ListenAndServe())
}
