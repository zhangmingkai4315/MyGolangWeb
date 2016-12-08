package routes

import (
	"data"
	"net/http"
	"templates"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	_, err := data.GetSession(w, r)
	render_data := make(map[string]interface{})
	if err == nil {
		render_data["Auth"] = true
		templates.RenderTemplate(w, "404", render_data)

	} else {
		render_data["Auth"] = false
		templates.RenderTemplate(w, "404", render_data)
	}
}
