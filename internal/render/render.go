package render

import (
	"bytes"
	"fmt"
	"github.com/justinas/nosurf"
	"github.com/natib21/bookings/internal/config"
	"github.com/natib21/bookings/internal/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, req *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(req)
	return td
}

func RenderTemplate(res http.ResponseWriter, req *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Couid not get template from template cache")
	}
	buf := new(bytes.Buffer)

	td = AddDefaultData(td, req)

	err := t.Execute(buf, td)
	if err != nil {
		log.Printf("Error executing template %s: %v", tmpl, err)
		return
	}

	_, err = buf.WriteTo(res)
	if err != nil {
		fmt.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
