package handlers

import (
	"github.com/natib21/bookings/pkg/config"
	"github.com/natib21/bookings/pkg/models"
	"github.com/natib21/bookings/pkg/render"
	"log"
	"net/http"
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
func NewHandlers(r *Repository) {
	Repo = r
}
func (m *Repository) Home(res http.ResponseWriter, req *http.Request) {
	remoteIp := req.RemoteAddr
	m.App.Session.Put(req.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(res, "home.page.gohtml", &models.TemplateData{})
}
func (m *Repository) About(res http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello,again"
	log.Printf("StringMap in About handler: %+v", stringMap)

	remoteIp := m.App.Session.GetString(req.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(res, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
