package handlers

import (
	"github.com/m-r-ebrahimi-dev/golang-simple-web-app/pkg/render"
	"net/http"
)

// HomePage is about index page
func HomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is intro page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
