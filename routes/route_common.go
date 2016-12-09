package routes

import (
	"net/http"
	"templates"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	render_data := statusCollection(w, r)
	templates.RenderTemplate(w, "404", render_data)
}
