package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

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

func addBreaks(oldpoem string) string {
	poem := []string{}
	for _, line := range strings.Split(oldpoem, "\n") {
		poem = append(poem, line+"<br/>")
	}
	return strings.Join(poem, "")
}

/*

<div class="row"> // i%5
	<div class="col"> //everytime
		<div class="card" style="width: 18rem;">
			<div class="card-header">
				Poem 1
			</div>
			<div class="card-body mb-4">
				<p class="card-text">
					{{index .StringMap "1"}}
				</p>
			</div>
		</div>
	</div>
</div>


*/

func (m *Repository) Poems(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	stringMap := make(map[string]string)
	stringMap["poem"] = ""
	rowCount := 0
	for i := 0; i < m.App.NumberOfPoems; i++ {
		poem := ""
		numStr := strconv.Itoa(i) // This allows 1 -> "1" not 1 -> smiley Face
		if i%5 == 0 {
			poem += "<div class=\"row\">\n"
			rowCount++
		}
		poem += "<div class=\"col\">\n<div class=\"card\" style=\"width: 18rem;\">\n<div class=\"card-header\">"
		poem += "Poem " + numStr
		poem += "</div>\n<div class=\"card-body mb-4\">\n<p class=\"card-text\">\n"
		filePath := "../../static/poems/" + numStr + ".txt"
		dat, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
		}
		poem += addBreaks(string(dat))

		poem += "</p></div></div></div>"
		if i%5-1 == 0 || i+1 == m.App.NumberOfPoems {
			poem += "</div>"
		}
		stringMap["poem"] += poem
	}

	render.RenderTemplate(w, "poems.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About is the about page handler
func (m *Repository) Test(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "You've been Chazzzed"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "test.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
