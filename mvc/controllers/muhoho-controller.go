package controllers

import (
	"net/http"

	"github.com/Greek-Academy/software-architecture/mvc/exceptions"
	"github.com/Greek-Academy/software-architecture/mvc/models"
	"github.com/Greek-Academy/software-architecture/mvc/views"
)

func MuhohoController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, exceptions.MethodNotAllowed().Error(), http.StatusMethodNotAllowed)
		return
	}

	// リクエストの受け取り
	eye := r.URL.Query().Get("eye")
	mouth := r.URL.Query().Get("mouth")

	// ビジネスロジック（Model）
	muhoho := models.NewMuhoho(&w, eye, mouth)

	// レンダリング（View）
	v := views.NewMuhoho(&w, muhoho)
	if err := v.Render(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
