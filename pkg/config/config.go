package config

import (
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application configuration
type AppConfig struct {
	UseCache       bool
	TemplateCache  map[string]*template.Template
	InProduction   bool
	NumberOfPoems  int
	NumberOfBlogs  int
	Session        *scs.SessionManager
	FractalWidth   int
	FractalHeight  int
	EscapeDistance int
}
