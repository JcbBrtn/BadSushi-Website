package handlers

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"net/http"
	"sync"

	"github.com/jcbbrtn/BadSushi/pkg/config"
	"github.com/jcbbrtn/BadSushi/pkg/models"
	"github.com/jcbbrtn/BadSushi/pkg/render"
)

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
		//n = Cos_z0(n, c)
		distance = Magnitude(n)
		count++
	}
	//Insert Coloring algorithm here

	color := color.RGBA{R: uint8(count), G: uint8(count), B: uint8(count), A: 255}

	//Set the pixle and end the thread
	img.SetRGBA(int(x), int(y), color)
}

func Mandlebrot(x complex128, c complex128) complex128 {
	return (x * x) + c
}

func Magnitude(x complex128) float64 {
	return math.Sqrt(real(x)*real(x) + imag(x)*imag(x))
}

func MyFunc(n complex128) complex128 {
	x := real(n)
	y := imag(n)
	return complex(-1*math.Sin(y)*math.Sinh(x), 1/(math.Cos(y)*math.Cosh(x)))
}

func Cos_z0(z complex128, z0 complex128) complex128 {
	x := real(z)
	y := imag(z)

	a := (math.Exp(x) + math.Exp(-x)) / 2
	b := (math.Exp(y) + math.Exp(-y)) / 2
	return complex(a, b) + z0
}
