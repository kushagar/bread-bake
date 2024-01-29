package main

import (
	"bread-bake/pkg/config"
	handler "bread-bake/pkg/handler"
	"bread-bake/pkg/render"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig
var session *scs.SessionManager

const portNumber = ":8080"

// main is the main function
func main() {

	app = config.AppConfig{}
	tc, err := render.CreateTemplateCache()

	//session options
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // true in production

	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	repo := handler.NewRepo(&app)
	handler.NewHandler(repo)
	app.Session = session
	app.TemplateCache = tc
	app.UseCache = false
	app.InProduction = false
	render.SetConfig(&app)
	router := routes(&app)
	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	srv := http.Server{
		Addr:    portNumber,
		Handler: router,
	}
	_ = srv.ListenAndServe()
}
