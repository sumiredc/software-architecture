package web

import (
	"net/http"

	"github.com/Greek-Academy/software-architecture/clean-architecture/drivers/web/route_bind"
)

func Router() {
	http.HandleFunc("/muhoho", route_bind.MuhohoRouteBind)
}
