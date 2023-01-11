package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/m-r-ebrahimi-dev/golang-simple-web-app/pkg/config"
	"github.com/m-r-ebrahimi-dev/golang-simple-web-app/pkg/handlers"
	"github.com/m-r-ebrahimi-dev/golang-simple-web-app/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// todo: change InProduction to true
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error creating template cache: ", err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		return
	}
}
