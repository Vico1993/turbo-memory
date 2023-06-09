package bot

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type botCmd interface {
	Command() string
	Exec(Bot *tgbotapi.BotAPI, message *tgbotapi.Message) error
}

// Return the list of Command
var CmdList []botCmd

func Init() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	// Load commands
	CmdList = append(CmdList, newAddCmd())

	fmt.Println("Bot command loaded!")

	for update := range updates {
		handleUpdates(update, bot)
	}
}
