package routes

import (
	"data"
	"net/http"
	"templates"
)

// index function will handle the '/' query, return the threads data for users
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//check if the user is a anthenticted user.

	threads, err := data.Threads()
	render_data:=statusCollection(w,r)
	if err == nil {
		render_data["Threads"] = threads
		templates.RenderTemplate(w, "home", render_data)
	} else {
		//should redirect to a server fail page.
		templates.RenderTemplate(w, "404", render_data)
	}
}
