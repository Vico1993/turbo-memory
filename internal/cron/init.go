package cron

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-co-op/gocron"
)

func Init() {
	s := gocron.NewScheduler(time.UTC)

	// Check every day how many you spent
	_, err := s.Every(1).Day().At("21:00").Do(func() {
		telegramPostMessage("Don't forget to log your expenses")
	})
	if err != nil {
		fmt.Println("Couldn't run the Check Message")
	}

	// Be polite every morning
	_, err = s.Every(1).Day().At("08:00").Do(func() {
		quotes := []string{
			`Believe you can and you're halfway there. - Theodore Roosevelt`,
			`You are never too old to set another goal or to dream a new dream. - C.S. Lewis`,
			`Success is not final, failure is not fatal: it is the courage to continue that counts. - Winston Churchill`,
			`Every morning we are born again. What we do today matters most. - Buddha`,
			`The only way to do great work is to love what you do. - Steve Jobs`,
			`The future belongs to those who believe in the beauty of their dreams. - Eleanor Roosevelt`,
			`Don't watch the clock; do what it does. Keep going. - Sam Levenson`,
			`You miss 100% of the shots you don't take. -  Wayne Gretzky`,
			`Believe in yourself and all that you are. Know that there is something inside you that is greater than any obstacle. - Christian D. Larson`,
			`The greatest glory in living lies not in never falling, but in rising every time we fall. - Nelson Mandela`,
		}

		telegramPostMessage("Good morning Victor!")
		telegramPostMessage(quotes[rand.Intn(len(quotes))])
	})
	if err != nil {
		fmt.Println("Couldn't run the Check Message")
	}

	// Start executing cron Async
	s.StartAsync()

	fmt.Println("Cron job initied!")
}
