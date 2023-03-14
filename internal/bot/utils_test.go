package bot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildMessage(t *testing.T) {
	res := buildMessage(1234, "Hello", 12)

	assert.Equal(t, "HTML", res.ParseMode, "Parse Mode should be HTML")
	assert.Equal(t, 12, res.ReplyToMessageID, "Reply Message should be correctly set")
	assert.Equal(t, int64(1234), res.ChatID, "Message ID should be correctly set")
	assert.Equal(t, "Hello", res.Text, "Text should be correctly set")
}
