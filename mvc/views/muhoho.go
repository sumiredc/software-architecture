package views

import (
	"net/http"
	"text/template"

	"github.com/Greek-Academy/software-architecture/mvc/models"
	"github.com/Greek-Academy/software-architecture/mvc/responses"
)

const (
	muhoho_template = "templates/muhoho.html"
)

type Muhoho struct {
	writer   *http.ResponseWriter
	template string
	muhoho   *models.Muhoho
}

func NewMuhoho(w *http.ResponseWriter, m *models.Muhoho) *Muhoho {
	return &Muhoho{
		writer:   w,
		template: muhoho_template,
		muhoho:   m,
	}
}

func (v *Muhoho) Render() error {
	t, err := template.ParseFiles(v.template)
	if err != nil {
		return err
	}

	response := responses.NewFaceResponse(v.muhoho.Face())
	return t.Execute(*v.writer, response)
}
