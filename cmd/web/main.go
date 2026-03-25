package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/I-Maged/00-golang-first-server/pkg/config"
	"github.com/I-Maged/00-golang-first-server/pkg/handlers"
	"github.com/I-Maged/00-golang-first-server/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHnadlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting server on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
