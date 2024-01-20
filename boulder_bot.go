package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	log.Printf("beep boop!")
	log.Printf(os.Getenv("API_TOKEN"))
}
