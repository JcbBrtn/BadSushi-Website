package handlers

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"

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
			<div class="row"> // i%4
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

func (m *Repository) Fractal(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "fractal.page.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) Fractal_Render(w http.ResponseWriter, r *http.Request) {
	imageWidth := m.App.FractalWidth
	imageHeight := m.App.FractalHeight
	var wg sync.WaitGroup
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	for x := 0; x < imageWidth; x++ {
		for y := 0; y < imageHeight; y++ {
			wg.Add(1)
			go setPixleColor(float64(x), float64(y), img, &wg, *m.App)
		}
	}
	wg.Wait()
	png.Encode(w, img)
}

func setPixleColor(x float64, y float64, img *image.RGBA, wg *sync.WaitGroup, app config.AppConfig) {
	defer wg.Done()
	//Get the coordinate this pixle represents
	a := (float64(app.RealRange)/float64(app.FractalHeight))*x - (float64(app.RealRange) / 2)
	b := (float64(app.CompRange)/float64(app.FractalWidth))*y - (float64(app.CompRange) / 2)
	c := complex(a, b) //first instance of n, this can be reffered to as z0
	n := c
	count := 0
	distance := 0.0
	for count < 255 && distance < float64(app.EscapeDistance) {
		//n = Mandlebrot(n, c)
		n = MyFunc(n)
		distance = Magnitude(n)
		count++
	}
	//Insert Coloring algorithm here

	color := color.RGBA{R: uint8(count), G: uint8(count), B: uint8(count), A: 255}

	//Set the pixle and end the thread
	img.SetRGBA(int(x), int(y), color)
}
