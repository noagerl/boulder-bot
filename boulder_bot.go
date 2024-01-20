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
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tg.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {
		// Telegram can send many types of updates depending on what your Bot
		// is up to. We only want to look at messages for now, so we can
		// discard any other updates.
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		// msg := tg.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			send_msg(bot, update.Message.Chat.ID, "Currently only /status is supported.")
		case "status":
			send_msg(bot, update.Message.Chat.ID, "Will create next poll at 2024-01-20 18:00")
		case "boulderpoll":
			start_boulder_poll(bot, update.Message.Chat.ID)
		}

		// Okay, we're sending our message off! We don't care about the message
		// we just sent, so we'll discard it.
		// if _, err := bot.Send(msg); err != nil {
		// 	// Note that panics are a bad way to handle errors. Telegram can
		// 	// have service outages or network errors, you should retry sending
		// 	// messages or more gracefully handle failures.
		// 	panic(err)
		// }
	}
}

func send_msg(bot *tg.BotAPI, chat_id int64, text string) {
	// Okay, we're sending our message off! We don't care about the message
	// we just sent, so we'll discard it.
	msg := tg.NewMessage(chat_id, text)
	if _, err := bot.Send(msg); err != nil {
		// Note that panics are a bad way to handle errors. Telegram can
		// have service outages or network errors, you should retry sending
		// messages or more gracefully handle failures.
		panic(err)
	}
}

func start_boulder_poll(bot *tg.BotAPI, chat_id int64) {
	msg := tg.NewPoll(chat_id, "What's the largest object in the universe?", "The Sun", "Sagittarius A*", "PSR J0952–0607", "Your mom")
	if _, err := bot.Send(msg); err != nil {
		// Note that panics are a bad way to handle errors. Telegram can
		// have service outages or network errors, you should retry sending
		// messages or more gracefully handle failures.
		panic(err)
	}
}
