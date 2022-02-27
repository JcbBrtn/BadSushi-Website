package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jcbbrtn/BadSushi/pkg/config"
	"github.com/jcbbrtn/BadSushi/pkg/handlers"
	"github.com/jcbbrtn/BadSushi/pkg/render"
)

//Uncomment this for production environment
var portNumber = ":" + os.Getenv("PORT")

//Web address will be localhost:8080
//This should only be used in Test environment
//Make sure to comment this line when pushing to production
//var portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//Set the Config for the fractal generator
	app.FractalHeight = 1080
	app.FractalWidth = 1080
	app.EscapeDistance = 100
	app.InProduction = true
	app.RealRange = 10
	app.CompRange = 10

	//Read in all the files in the poems folder
	files, err := ioutil.ReadDir("static/poems")
	if err != nil {
		panic(err)
	}
	app.NumberOfPoems = len(files)
	fmt.Println("Number of poems loaded: ", app.NumberOfPoems)

	//Read in all files in the blogs folder
	files, err = ioutil.ReadDir("static/blogs")
	if err != nil {
		panic(err)
	}
	app.NumberOfBlogs = len(files)
	fmt.Println("Number of blogs loaded: ", app.NumberOfBlogs)

	//Use the middleware to create a new session
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

	//Start the server
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
