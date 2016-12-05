package templates

import (
	"html/template"
	"log"
)

var public_tmpl_files = []string{
	"templates/layout.html",
	"templates/public.navbar.html",
	"templates/index.html",
}
var private_tmpl_files = []string{
	"templates/layout.html",
	"templates/private.navbar.html",
	"templates/index.html",
}

//define the templates for differents users
//private_templates : authenticated user will using it
//public_templates : unauthenticated users will using it
var Public_templates, Private_templates *template.Template

func InitTemplates() {
	//Must function will throw panic when parse failure. That's must done works before start the app!
	Private_templates = template.Must(template.ParseFiles(public_tmpl_files...))
	Public_templates = template.Must(template.ParseFiles(private_tmpl_files...))
	log.Println("Load templates to memory done!")
}
