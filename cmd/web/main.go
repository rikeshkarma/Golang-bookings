package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/rikeshkarma/Golang-bookings/pkg/config"
	"github.com/rikeshkarma/Golang-bookings/pkg/handlers"
	"github.com/rikeshkarma/Golang-bookings/pkg/render"
)

// Divide page handler

var app config.AppConfig
var session *scs.SessionManager

func main() {
	

	app.InProd = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProd
	app.Session = session


	tc, err := render.CreateTemplateCache()
	if err !=nil {
		log.Fatal("Cannot get template cache")
	}

	app.UseCache = false
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server {
		Addr: ":8080",
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	log.Fatal(err)
}