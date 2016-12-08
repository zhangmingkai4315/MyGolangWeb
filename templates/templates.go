package templates

import (
	"fmt"
	"github.com/oxtoacart/bpool"
	"github.com/spf13/viper"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templates map[string]*template.Template
var bufpool *bpool.BufferPool

func init() {
	bufpool = bpool.NewBufferPool(64)
}
func InitTemplates() {
	//Must function will throw panic when parse failure. That's must done works before start the app!
	//Private_templates = template.Must(template.ParseFiles(private_tmpl_files...))
	//Public_templates = template.Must(template.ParseFiles(public_tmpl_files...))
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	templatesDir := viper.GetString("templates")
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		log.Fatalln(err)
	}
	base := layouts[0]
	pages, err := filepath.Glob(templatesDir + "/pages/*.html")
	if err != nil {
		log.Fatalln(err)
	}
	includes, err := filepath.Glob(templatesDir + "/includes/*/*.html")
	if err != nil {
		log.Fatalln(err)
	}
	for _, page := range pages {
		files := append(includes, base, page)
		filename := filepath.Base(page)
		extension := filepath.Ext(page)
		templates[filename[0:len(filename)-len(extension)]] = template.Must(template.ParseFiles(files...))
	}
	log.Println("Load templates to memory done")
}

func RenderTemplate(w http.ResponseWriter, name string, data map[string]interface{}) error {
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("Template %s does not exist", name)
	}
	// Create a buffer to temporarily write to and check if any errors were encounted.
	buf := bufpool.Get()
	defer bufpool.Put(buf)

	err := tmpl.ExecuteTemplate(buf, "base", data)
	if err != nil {
		log.Println(err)
		return err
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)

	return nil
}
