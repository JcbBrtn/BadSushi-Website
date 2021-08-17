package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jcbbrtn/BadSushi/pkg/config"
	"github.com/jcbbrtn/BadSushi/pkg/handlers"
	"github.com/jcbbrtn/BadSushi/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.FractalHeight = 1080
	app.FractalWidth = 1080
	app.EscapeDistance = 100
	app.InProduction = false

	files, err := ioutil.ReadDir("../../static/poems")
	if err != nil {
		panic(err)
	}
	app.NumberOfPoems = len(files)
	fmt.Println("Number of poems loaded: ", app.NumberOfPoems)

	files, err = ioutil.ReadDir("../../static/blogs")
	if err != nil {
		panic(err)
	}
	app.NumberOfBlogs = len(files)
	fmt.Println("Number of blogs loaded: ", app.NumberOfBlogs)

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")

	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
