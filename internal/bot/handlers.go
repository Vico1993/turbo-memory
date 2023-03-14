package bot

import (
	"fmt"
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

	fmt.Println("Talking with(id):", update.Message.Chat.ID)
	fmt.Println("Talking with(firstName):", update.Message.Chat.FirstName)
	fmt.Println("Talking with(lastName):", update.Message.Chat.LastName)
	fmt.Println("Talking with(userName):", update.Message.Chat.UserName)
	fmt.Println("Type", update.FromChat().Type)
}
