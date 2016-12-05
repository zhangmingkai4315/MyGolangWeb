package routes

import (
	"data"
	"fmt"
	"log"
	"net/http"
	"templates"
)

// index function will handle the '/' query, return the threads data for users
func IndexHandler(w http.ResponseWriter, r *http.Request) {
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

func ErrHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Someone's request is broken!", r.URL.Path)
	fmt.Fprintln(w, "This is a error page!")
}
