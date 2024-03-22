package controllers

import (
	"net/http"

	"github.com/Greek-Academy/software-architecture/mvc/exceptions"
	"github.com/Greek-Academy/software-architecture/mvc/views"
)

func IndexController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, exceptions.MethodNotAllowed().Error(), http.StatusMethodNotAllowed)
		return
	}

	// レンダリング（View）
	v := views.NewIndex(&w)
	if err := v.Render(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
