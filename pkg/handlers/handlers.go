package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/rafaelsfaria/bookings/pkg/config"
	"github.com/rafaelsfaria/bookings/pkg/models"
	"github.com/rafaelsfaria/bookings/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

// handlers' Repository
var Repo *Repository

// creates new repo
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// set repositor for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the about home handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello world"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func divideNumber(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("cannot divide by 0")
	}
	result = x / y
	return result, nil
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var dividend float32 = 10.0
	var divisor float32 = 0
	result, err := divideNumber(dividend, divisor)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	str := fmt.Sprintf("%f divided by %f is %f", dividend, divisor, result)
	fmt.Fprint(w, str)
}
