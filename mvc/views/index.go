package views

import (
	"net/http"
	"text/template"
)

const (
	index_template = "templates/index.html"
)

type Index struct {
	writer   *http.ResponseWriter
	template string
}

func NewIndex(w *http.ResponseWriter) *Index {
	return &Index{
		writer:   w,
		template: index_template,
	}
}

func (v *Index) Render() error {
	t, err := template.ParseFiles(v.template)
	if err != nil {
		return err
	}

	return t.Execute(*v.writer, nil)
}
