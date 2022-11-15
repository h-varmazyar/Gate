package tgBotApi

import (
	botAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Keyboard struct {
	UseInline   bool
	UseKeyboard bool
	Inline      *botAPI.InlineKeyboardMarkup
	Keyboard    *botAPI.ReplyKeyboardMarkup
}
