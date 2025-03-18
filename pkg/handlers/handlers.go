package handlers

import (
	"net/http"

	"github.com/irsyadfata99/go-sec-app/pkg/config"
	"github.com/irsyadfata99/go-sec-app/pkg/models"
	"github.com/irsyadfata99/go-sec-app/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// Reservation is the handler for the reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {    
    render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}


// Generals is the handler for the generals page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.page.tmpl", &models.TemplateData{})
}

// Majorsuite is the handler for the majorsuite page
func (m *Repository) MajorSuite(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.page.tmpl", &models.TemplateData{})
}

// Fix this functionfunc (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
		// Change this from "search-availbility.page.tmpl" to "search-availability.page.tmpl"
		render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
	}