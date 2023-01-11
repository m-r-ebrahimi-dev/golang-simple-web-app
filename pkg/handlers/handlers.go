package handlers

import (
	"github.com/m-r-ebrahimi-dev/golang-simple-web-app/pkg/config"
	"github.com/m-r-ebrahimi-dev/golang-simple-web-app/pkg/models"
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
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is intro page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello, there"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
