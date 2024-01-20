package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
)

func main() {
	scheduler, err := gocron.NewScheduler()

	if err != nil {
		log.Fatal(err)
	}

	defer func() { _ = scheduler.Shutdown() }()
	gocron.AfterJobRunsWithError(func(jobID uuid.UUID, jobName string, err error) {
		//TODO send error into some admin group channel
	})

	//schedule := gocron.WeeklyJob(1, gocron.NewWeekdays(time.Sunday), gocron.NewAtTimes(gocron.NewAtTime(11, 0, 0)))
	schedule := gocron.DurationJob(time.Second * 3)
	task := gocron.NewTask(func() {
		fmt.Println("Hi!")
	})

	j, err := scheduler.NewJob(schedule, task)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Added job %s with ID %s", j.Name(), j.ID())
	log.Printf("Starting scheduler ...")
	scheduler.Start()
	log.Printf("Scheduler started!")
	select {}
}
