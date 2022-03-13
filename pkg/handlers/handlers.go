package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/jcbbrtn/BadSushi/pkg/config"
	"github.com/jcbbrtn/BadSushi/pkg/models"
	"github.com/jcbbrtn/BadSushi/pkg/render"
)

var Repo *Repository

//Repository is the type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handlering
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

//A Simple Game of catch phrase where the device gives the words.
func (m *Repository) Catch_Phrase(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "cp.page.html", &models.TemplateData{})
}

//Here is where I can test things - mechanics wild techniques that I don't have any solid implimentation for yet.
func (m *Repository) Lab(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	stringMap["ip"] = m.App.Session.GetString(r.Context(), "remote_ip")

	render.RenderTemplate(w, "lab.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Todo_List(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "todolist.page.html", &models.TemplateData{})
}

func (m *Repository) Tell_Me_Im_Right(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	valids := make([]string, 0)
	valids = append(valids,
		"You're right!",
		"That's correct!",
		"You've hit the nail on the head!",
		"Got it!",
		"I couldn't have said it any better myself.",
		"That's spot on.",
		"Absolutley.",
		"You could say so.",
		"There's nothing to add on that.",
		"You're dead right.",
		"Exactly.",
	)

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	stringMap["statement"] = valids[rand.Intn(len(valids))]

	render.RenderTemplate(w, "tmir.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
