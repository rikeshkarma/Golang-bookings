package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/rikeshkarma/Golang-bookings/pkg/config"
	"github.com/rikeshkarma/Golang-bookings/pkg/models"
)

var functions = template.FuncMap {

}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders templates
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	tc := app.TemplateCache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Couldn't get template from tempalte cache")
	}
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)
	buf.WriteTo(w)
}

func CreateTemplateCache() (map[string]*template.Template, error){

	myCache := make(map[string]*template.Template)

	pages, err := filepath.Glob("templates/*.page.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts , err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		matches, err := filepath.Glob("templates/*.layout.html")
		if err != nil {
			return nil, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("templates/*.layout.html")
			if err != nil {
				return nil, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}