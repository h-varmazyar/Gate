package tgBotApi

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	errorext "github.com/h-varmazyar/Gate/pkg/errors"
	"google.golang.org/grpc/codes"
	"regexp"
)

type Message interface {
	send() error
}

type BaseMessage struct {
	ChatID           int64
	ReplyToMessageID int
}

func SendMessage(ctx context.Context, msg Message) error {
	if !running {
		return errorext.New(ctx, codes.Aborted)
	}

	if err := msg.send(); err != nil {
		return err
	}
	return nil
}

type textMessage struct {
	BaseMessage
	Keyboard *Keyboard
	Text     string
}

func NewTextMessage(chatID int64, replyTo int, content string, keyboard *Keyboard) *textMessage {
	return &textMessage{
		BaseMessage: BaseMessage{
			ChatID:           chatID,
			ReplyToMessageID: replyTo,
		},
		Keyboard: keyboard,
		Text:     content,
	}
}

func (msg *textMessage) send() error {
	tgMessage := tgbotapi.NewMessage(msg.ChatID, msg.Text)
	if msg.ReplyToMessageID > 0 {
		tgMessage.ReplyToMessageID = msg.ReplyToMessageID
	}
	if msg.Keyboard != nil {
		if msg.Keyboard.UseInline {
			tgMessage.ReplyMarkup = msg.Keyboard.Inline
		} else if msg.Keyboard.UseKeyboard {
			tgMessage.ReplyMarkup = msg.Keyboard.Keyboard
		}
	}
	r, _ := regexp.Compile("[[a-zA-Z\\d]*]\\(tg:\\/\\/user\\?id=\\d*\\)")
	if r.MatchString(msg.Text) {
		tgMessage.ParseMode = "Markdown"
	}
	if _, err = bot.Send(tgMessage); err != nil {
		return err
	}
	return nil
}
