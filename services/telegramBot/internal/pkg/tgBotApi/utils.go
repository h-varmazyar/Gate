package tgBotApi

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CallbackData struct {
	Query string
	Data  interface{}
}

func NewCallbackDataButton(query string, data interface{}) tgbotapi.InlineKeyboardButton {
	callbackData := &CallbackData{
		Query: query,
		Data:  data,
	}

	dataBytes, _ := json.Marshal(callbackData)
	return tgbotapi.NewInlineKeyboardButtonData(query, string(dataBytes))
}

func ParseCallbackData(callbackData string) (query string, data interface{}, err error) {
	tmp := new(CallbackData)
	if err = json.Unmarshal([]byte(callbackData), tmp); err != nil {
		return "", nil, err
	}
	return tmp.Query, tmp.Data, nil
}
