package main

import (
	"net/http"

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
	mux.Get("/test", handlers.Repo.Test)
	mux.Get("/poems", handlers.Repo.Poems)
	mux.Get("/blog", handlers.Repo.Log)

	//Find way to create a seperate page for each blog.

	return mux
}
