package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/jcbbrtn/BadSushi/pkg/config"
	"github.com/jcbbrtn/BadSushi/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
		//fmt.Println("Got Template Cache")
	} else {
		var err error
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatal(err)
		}
	}

	t, ok := tc[tmpl]
	if !ok {
		fmt.Println("error finding template from cache")
		log.Fatal()
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error parsing template: ", err)
	}
}

//CreateTemplateCache creates a cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("templates/*.page.html")

	if err != nil {
		fmt.Println("Error 0")
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println("Error 1 With name: ", name)
			return myCache, err
		}

		matches, err := filepath.Glob("templates/*.layout.html")
		if err != nil {
			fmt.Println("Error 2")
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("templates/*.layout.html")
			if err != nil {
				fmt.Println("Error 3")
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
