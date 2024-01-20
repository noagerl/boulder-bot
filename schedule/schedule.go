package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func main() {
	s, err := gocron.NewScheduler()

	if err != nil {
		// handle error
	}

	defer func() { _ = s.Shutdown() }()

	//schedule := gocron.WeeklyJob(1, gocron.NewWeekdays(time.Sunday), gocron.NewAtTimes(gocron.NewAtTime(11, 0, 0)))
	schedule := gocron.DurationJob(time.Second * 3)
	task := gocron.NewTask(func() {
		fmt.Println("Hi!")
	})

	j, err := s.NewJob(schedule, task)
	if err != nil {
		// handle error
	}

	log.Printf("Added job %s with ID %s", j.Name(), j.ID())
	log.Printf("Starting scheduler ...")
	s.Start()
	log.Printf("Scheduler started!")
	select {}
}
