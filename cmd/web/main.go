package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rafaelsfaria/bookings/pkg/config"
	"github.com/rafaelsfaria/bookings/pkg/handlers"
	"github.com/rafaelsfaria/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var session *scs.SessionManager

func main() {
	var app config.AppConfig

	// set to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Println("Listening on port", portNumber)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
