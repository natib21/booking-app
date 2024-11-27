package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/natib21/bookings/internal/config"
	"github.com/natib21/bookings/internal/forms"
	"github.com/natib21/bookings/internal/models"
	"github.com/natib21/bookings/internal/render"
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
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(res, req, "make-reservation.page.gohtml", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})

}
func (m *Repository) PostReservation(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	reservation := models.Reservation{
		FirstName: req.Form.Get("first_name"),
		LastName:  req.Form.Get("last_name"),
		Phone:     req.Form.Get("phone"),
		Email:     req.Form.Get("email"),
	}

	form := forms.New(req.PostForm)

	//form.Has("first_name", req)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, req)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(res, req, "make-reservation.page.gohtml", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return

	}
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
