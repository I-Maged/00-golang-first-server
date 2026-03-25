package main

import (
	"net/http"

	"github.com/I-Maged/00-golang-first-server/pkg/config"
	"github.com/I-Maged/00-golang-first-server/pkg/handlers"
	"github.com/bmizerany/pat"
)

func routes(a *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
