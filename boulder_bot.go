package main

import (
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/joho/godotenv"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	godotenv.Load()

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

	scheduler, err := gocron.NewScheduler()

	if err != nil {
		log.Fatal(err)
	} else {
		schedule := gocron.WeeklyJob(1, gocron.NewWeekdays(time.Sunday), gocron.NewAtTimes(gocron.NewAtTime(11, 0, 0)))
		task := gocron.NewTask(func() {
			// -1002119201796 is the chat id of our boulder_bot_test chat
			// this needs to be somehow registered, stored and loaded on startup
			startBoulderPoll(bot, -1002119201796)
		})
		j, err := scheduler.NewJob(schedule, task)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Added job %s with ID %s", j.Name(), j.ID())
		log.Printf("Starting scheduler ...")
		scheduler.Start()
		log.Printf("Scheduler started!")
	}

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
			sendMsg(bot, update.Message.Chat.ID, "Currently only /status and /boulderpoll is supported.")
		case "status":
			sendMsg(bot, update.Message.Chat.ID, "Will create next poll at 2024-01-20 18:00")
		case "boulderpoll":
			startBoulderPoll(bot, update.Message.Chat.ID)
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

func sendMsg(bot *tg.BotAPI, chatId int64, text string) {
	log.Printf("ChatID:%d", chatId)
	// Okay, we're sending our message off! We don't care about the message
	// we just sent, so we'll discard it.
	msg := tg.NewMessage(chatId, text)
	if _, err := bot.Send(msg); err != nil {
		// Note that panics are a bad way to handle errors. Telegram can
		// have service outages or network errors, you should retry sending
		// messages or more gracefully handle failures.
		panic(err)
	}
}

func startBoulderPoll(bot *tg.BotAPI, chatId int64) {
	dayPoll := tg.SendPollConfig{
		BaseChat: tg.BaseChat{
			ChatID: chatId,
		},
		Question:              "Ich will bouldern am",
		Options:               []string{"Montag", "Dienstag", "Mittwoch", "Donnerstag", "Freitag", "Samstag", "Sonntag"},
		IsAnonymous:           false,
		AllowsMultipleAnswers: true,
	}

	locationPoll := tg.SendPollConfig{
		BaseChat: tg.BaseChat{
			ChatID: chatId,
		},
		Question:              "Ich will bouldern in",
		Options:               []string{"Seestadt", "Wienerberg", "Hauptbahnhof", "Hannovermarkt", "Blockfabrik"},
		IsAnonymous:           false,
		AllowsMultipleAnswers: true,
	}

	if _, err := bot.Send(dayPoll); err != nil {
		// Note that panics are a bad way to handle errors. Telegram can
		// have service outages or network errors, you should retry sending
		// messages or more gracefully handle failures.
		panic(err)
	}
	if _, err := bot.Send(locationPoll); err != nil {
		// Note that panics are a bad way to handle errors. Telegram can
		// have service outages or network errors, you should retry sending
		// messages or more gracefully handle failures.
		panic(err)
	}
}
