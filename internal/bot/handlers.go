package bot

import (
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func shouldAct(update tgbotapi.Update) bool {
	chatUserId, _ := strconv.ParseInt(os.Getenv("TELEGRAM_USER_CHAT_ID"), 0, 64)

	// If it's not a Message or CallBackQuery
	// If it's not a Private chat with User Chat
	return !(update.Message == nil && update.CallbackQuery == nil) &&
		update.FromChat().Type == "private" &&
		update.Message.Chat.ID == chatUserId
}

func handleUpdates(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if !shouldAct(update) {
		// TODO: Log this why we don't act upon it
		return
	}

	if update.Message.IsCommand() {
		for _, cmd := range CmdList {
			if cmd.Command() == update.Message.Command() {
				cmd.Exec(bot, update.Message)
			}
		}
	}
}
