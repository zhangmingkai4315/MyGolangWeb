package routes

import (
	"data"
	"log"
	"net/http"
	"templates"
)

// index function will handle the '/' query, return the threads data for users
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		//check if the user is a anthenticted user.
		_, sessionErr := data.GetSession(w, r)
		render_data := map[string]interface{}{
			"Threads": threads,
		}
		if sessionErr == nil {
			render_data["Auth"] = true
		} else {
			render_data["Auth"] = false

		}
		templates.RenderTemplate(w, "home", render_data)
	} else {
		log.Println(err)
	}
}
