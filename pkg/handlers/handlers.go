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

// IndPost is a page for each induvidual blog post
func (m *Repository) IndPost(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	path := strings.Split(r.URL.Path, "/")
	filePath := "../../static/blogs/" + path[2] + ".html"
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	stringMap["post"] = string(dat)
	render.RenderTemplate(w, "post.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Log is the lading page for the main blog of BadSushi
func (m *Repository) Log(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	stringMap := make(map[string]string)
	stringMap["blogs"] = ""

	for i := 0; i < m.App.NumberOfBlogs; i++ {
		numStr := strconv.Itoa(i) // This allows 1 -> "1" not 1 -> smiley Face
		filePath := "../../static/blogs/" + numStr + ".html"
		dat, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
		}

		stringMap["blogs"] += string(dat)
		stringMap["blogs"] += "<hr/>"
	}

	render.RenderTemplate(w, "blog.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func addBreaks(oldpoem string) string {
	poem := []string{}
	for _, line := range strings.Split(oldpoem, "\n") {
		poem = append(poem, line+"<br/>")
	}
	return strings.Join(poem, "")
}

func (m *Repository) Poems(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	stringMap := make(map[string]string)
	stringMap["poem"] = ""

	for i := 0; i < m.App.NumberOfPoems; i++ {
		poem := ""
		numStr := strconv.Itoa(i) // This allows 1 -> "1" not 1 -> smiley Face
		filePath := "../../static/poems/" + numStr + ".txt"
		dat, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
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
								{{The Poems}}
							</p>
						</div>
					</div>
				</div>
			</div>
		*/

		if i%4 == 0 {
			poem += "<div class=\"row\">\n"
		}
		poem += "<div class=\"col\">\n<div class=\"card\" style=\"width: 18rem;\">\n<div class=\"card-header\">"
		poem += "Poem " + numStr
		poem += "</div>\n<div class=\"card-body text-secondary mb-4\">\n<p class=\"card-text\">\n"

		poem += addBreaks(string(dat))

		poem += "</p></div></div></div>"
		if i%4+1 == 0 || i+1 >= m.App.NumberOfPoems {
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
