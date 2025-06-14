package main

import (
	"github.com/joho/godotenv"
	"github.com/michaelwp/siloamteens/api/server"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	log.Fatal(server.Server(os.Getenv("PORT")))
}
