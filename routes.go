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
	mux.Use(SessionLoad)
	mux.Use(NoSurf)

	//Set the Endpoints
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/fractal", handlers.Repo.Fractal)
	mux.Get("/poems", handlers.Repo.Poems)
	mux.Get("/blog", handlers.Repo.Log)
	mux.Get("/fractal_render", handlers.Repo.Fractal_Render)
	mux.Get("/Catch_Phrase", handlers.Repo.Catch_Phrase)
	mux.Get("/lab", handlers.Repo.Lab)
	mux.Get("/lobby", handlers.Repo.SecretMissionLobby)
	mux.Get("/todo", handlers.Repo.Todo_List)
	mux.Get("/tell_me_im_right", handlers.Repo.Tell_Me_Im_Right)

	//create a seperate page for each blog.
	for i := 0; i < app.NumberOfBlogs; i++ {
		numStr := strconv.Itoa(i) // This allows 1 -> "1" not 1 -> smiley Face
		mux.Get("/blog/"+numStr, handlers.Repo.IndPost)
	}
	return mux
}
