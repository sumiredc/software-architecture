package main

import (
	"log"
	"os"

	"github.com/Greek-Academy/software-architecture/clean-architecture/drivers/web"
	"github.com/joho/godotenv"
)

func main() {
	// 環境変数の読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	web.Router()
	web.Server(os.Getenv("HOST"))
}
