package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jcbbrtn/BadSushi/pkg/config"
	"github.com/jcbbrtn/BadSushi/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//Use Middleware
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	//Set the Endpoints
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/fractal", handlers.Repo.Fractal)
	mux.Get("/poems", handlers.Repo.Poems)
	mux.Get("/blog", handlers.Repo.Log)
	mux.Get("/fractal_render", handlers.Repo.Fractal_Render)
	mux.Get("/worldBuilder", handlers.Repo.WorldBuilder)
	mux.Get("/worldBuilder/person", handlers.Repo.Person)
	mux.Get("/worldBuilder/faction", handlers.Repo.Faction)
	mux.Get("/worldBuilder/location", handlers.Repo.Location)
	mux.Get("/worldBuilder/kingdom", handlers.Repo.Kingdom)
	mux.Get("/worldBuilder/continent", handlers.Repo.Continent)
	mux.Get("/worldBuilder/items", handlers.Repo.Items)
	mux.Get("/worldBuilder/god", handlers.Repo.God)
	mux.Get("/worldBuilder/world", handlers.Repo.World)

	//create a seperate page for each blog.
	for i := 0; i < app.NumberOfBlogs; i++ {
		numStr := strconv.Itoa(i) // This allows 1 -> "1" not 1 -> smiley Face
		mux.Get("/blog/"+numStr, handlers.Repo.IndPost)
	}
	return mux
}
