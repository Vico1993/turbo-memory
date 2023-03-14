package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func buildMessage(chatID int64, text string, replyID int) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(
		chatID,
		text,
	)
	msg.ReplyToMessageID = replyID
	msg.ParseMode = tgbotapi.ModeHTML

	return msg
}
