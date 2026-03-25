package handlers

import (
	"net/http"

	"github.com/I-Maged/00-golang-first-server/pkg/config"
	"github.com/I-Maged/00-golang-first-server/pkg/models"
	"github.com/I-Maged/00-golang-first-server/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHnadlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateDate{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	render.RenderTemplate(w, "about.page.html", &models.TemplateDate{StringMap: stringMap})
}
