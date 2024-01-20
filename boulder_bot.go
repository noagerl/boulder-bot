package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	godotenv.Load()
	log.Printf("beep boop!")

	bot, err := tg.NewBotAPI(os.Getenv("API_TOKEN"))
}
