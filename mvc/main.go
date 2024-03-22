package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Greek-Academy/software-architecture/mvc/controllers"
	"github.com/joho/godotenv"
)

func main() {
	// 環境変数の読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	// ルーティング
	http.HandleFunc("/", controllers.IndexController)
	http.HandleFunc("/muhoho", controllers.MuhohoController)

	// Webサーバー 起動
	host := os.Getenv("HOST")
	err = http.ListenAndServe(host, nil)
	if err != nil {
		log.Fatal(err)
	}
}
