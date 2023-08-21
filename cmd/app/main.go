package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/lekht/bookwiki-grpc/internal/app"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	app.Run()
	time.Sleep(time.Hour)
}
