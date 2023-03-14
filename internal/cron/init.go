package cron

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func Init() {
	s := gocron.NewScheduler(time.UTC)

	// Check every day how many you spent
	_, err := s.Every(1).Day().At("18:00").Do(func() {
		telegramPostMessage("How many did you spent today?")
	})

	if err != nil {
		fmt.Println("Couldn't run the Check Message")
	}

	// Start executing cron Async
	s.StartAsync()
}
