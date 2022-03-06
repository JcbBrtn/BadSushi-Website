package handlers

import (
	"net/http"

	"github.com/jcbbrtn/BadSushi/pkg/models"
	"github.com/jcbbrtn/BadSushi/pkg/render"
)

type Lobby struct {
	Id string
}

func (m *Repository) SecretMissionLobby(w http.ResponseWriter, r *http.Request) {
	//HTML Page with two options, start new lobby or join a lobby
	render.RenderTemplate(w, "lobby.page.html", &models.TemplateData{})
}

func (m *Repository) SecretMissionGame(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "lab.page.html", &models.TemplateData{})
}

func (m *Repository) SmRefresh(w http.ResponseWriter, r *http.Request) {
	//API endpoint that updates the information on the lobby when the client calls this for an AJAX Update
}
