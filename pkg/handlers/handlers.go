package handlers

import (
	"github.com/m-r-ebrahimi-dev/golang-simple-web-app/pkg/config"
	"github.com/m-r-ebrahimi-dev/golang-simple-web-app/pkg/render"
	"net/http"
)

// Repo the repository used by handlers
var Repo *Repository

// Repository the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for new handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// HomePage is about index page
func (m *Repository) HomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is intro page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
