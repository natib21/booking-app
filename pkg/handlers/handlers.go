package handlers

import (
	"encoding/json"
	"fmt"
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
	render.RenderTemplate(res, req, "home.page.gohtml", &models.TemplateData{})
}
func (m *Repository) About(res http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello,again"
	log.Printf("StringMap in About handler: %+v", stringMap)

	remoteIp := m.App.Session.GetString(req.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(res, req, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
func (m *Repository) Reservation(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "make-reservation.page.gohtml", &models.TemplateData{})
}
func (m *Repository) Generals(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "generals.page.gohtml", &models.TemplateData{})
}

func (m *Repository) Majors(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "majours.page.gohtml", &models.TemplateData{})
}
func (m *Repository) Availibility(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "search-availability.page.gohtml", &models.TemplateData{})
}
func (m *Repository) PostAvailibility(res http.ResponseWriter, req *http.Request) {
	start := req.Form.Get("start")
	end := req.Form.Get("end")

	res.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

type jsonRes struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailibilityJson(res http.ResponseWriter, req *http.Request) {
	resp := jsonRes{
		Ok:      true,
		Message: "Available",
	}
	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}
	log.Println("Response:", string(out)) // Optional for debugging
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(out)
}
func (m *Repository) Contact(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "contact.page.gohtml", &models.TemplateData{})
}
