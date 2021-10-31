package handlers

import (
	"net/http"

	"github.com/jcbbrtn/BadSushi/pkg/config"
	"github.com/jcbbrtn/BadSushi/pkg/models"
	"github.com/jcbbrtn/BadSushi/pkg/render"
)

//THIS SECTION OF CODE HAS NOT BEEN IMPLEMENTED YET
type WorldBuild struct {
	AllPersons    []config.Person
	AllFactions   []config.Faction
	AllLocations  []config.Location
	AllKingdoms   []config.Kingdom
	AllContinents []config.Continent
	AllItems      []config.Items
	AllGods       []config.God
	AllWorlds     []config.World
}

//var build *WorldBuild

//Take new initial values and add them to the build
func (m *Repository) WorldBuilder(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "worldBuilderForm.page.html", &models.TemplateData{})
}

//Induvidual pages of each section of the build to make edits to

func (m *Repository) Person(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "worldBuilderForm.page.html", &models.TemplateData{})
}

func (m *Repository) Faction(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "worldBuilderForm.page.html", &models.TemplateData{})
}

func (m *Repository) Location(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "worldBuilderForm.page.html", &models.TemplateData{})
}

func (m *Repository) Kingdom(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "worldBuilderForm.page.html", &models.TemplateData{})
}

func (m *Repository) Continent(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "worldBuilderForm.page.html", &models.TemplateData{})
}

func (m *Repository) Items(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "worldBuilderForm.page.html", &models.TemplateData{})
}

func (m *Repository) God(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "worldBuilderForm.page.html", &models.TemplateData{})
}

func (m *Repository) World(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "worldBuilderForm.page.html", &models.TemplateData{})
}
