package bot

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

func TestShouldActWithMessageInGroup(t *testing.T) {
	res := shouldAct(tgbotapi.Update{
		Message: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{
				Type: "group",
			},
		},
	})

	assert.False(t, res, "Should be false as the type is group and not private")
}

func TestShouldActWithMessageOrCallBack(t *testing.T) {
	res := shouldAct(tgbotapi.Update{
		CallbackQuery: nil,
		Message:       nil,
		EditedMessage: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{
				Type: "private",
			},
		},
	})

	assert.False(t, res, "Should be false as Message and CallbackQuery are nil")
}

func TestShouldActWithIncorrectChatID(t *testing.T) {
	res := shouldAct(tgbotapi.Update{
		Message: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{
				Type: "private",
				ID:   312039401804,
			},
		},
	})

	assert.False(t, res, "Should be false as the chat ID is incorect")
}
