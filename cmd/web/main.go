package main

import (
	"github.com/m-r-ebrahimi-dev/golang-simple-web-app/pkg/config"
	"github.com/m-r-ebrahimi-dev/golang-simple-web-app/pkg/handlers"
	"github.com/m-r-ebrahimi-dev/golang-simple-web-app/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error creating template cache: ", err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.HomePage)
	http.HandleFunc("/about", handlers.Repo.About)

	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		return
	}
}
