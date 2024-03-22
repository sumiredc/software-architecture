package web

import (
	"log"
	"net/http"
)

func Server(host string) {
	// Webサーバー 起動
	err := http.ListenAndServe(host, nil)
	if err != nil {
		log.Fatal(err)
	}
}
