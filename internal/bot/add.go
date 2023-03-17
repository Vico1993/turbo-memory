package bot

import (
	"strconv"
	"strings"

	"github.com/Vico1993/turbo-memory/internal/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type addCmd struct{}

func newAddCmd() *addCmd {
	return &addCmd{}
}

func (c *addCmd) Command() string {
	return "add"
}

func (c *addCmd) Exec(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	params := strings.Split(message.CommandArguments(), " ")

	// Parse information
	value, err := strconv.ParseFloat(params[0], 64)
	if err != nil {
		// couldn't load the value
		return err
	}

	row := database.NewRow(value, params[1:]...)
	row.Save()

	_, err = bot.Send(
		buildMessage(
			message.Chat.ID,
			"Thank you, it's added",
			message.MessageID,
		),
	)
	if err != nil {
		return err
	}

	return nil
}
