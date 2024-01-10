package tgBotApi

import (
	"context"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	errorext "github.com/h-varmazyar/Gate/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"sync"
	"time"
)

var (
	err        error
	running    bool
	bot        *tgbotapi.BotAPI
	updates    tgbotapi.UpdatesChannel
	lock       = &sync.Mutex{}
	h          Handlers
	botConfigs *Configs
)

func Run(customHandlers Handlers, configs *Configs) error {
	lock.Lock()
	defer lock.Unlock()

	if customHandlers == nil {
		return errors.New("please specify handler")
	} else {
		h = customHandlers
	}

	botConfigs = configs

	bot, err = tgbotapi.NewBotAPI(botConfigs.Token)
	if err != nil {
		log.WithError(err).Error("failed to create new bot")
		return err
	}

	bot.Debug = configs.DebugMode

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates = bot.GetUpdatesChan(u)
	go func() {
		for update := range updates {
			handleUpdate(update)
		}
	}()
	running = true
	return nil
}

func Stop(ctx context.Context) error {
	lock.Lock()
	defer lock.Unlock()

	if bot != nil {
		bot = nil
	} else {
		return errorext.New(ctx, codes.FailedPrecondition).AddDetailF("bot stopped before")
	}

	if updates != nil {
		updates.Clear()
	} else {
		return errorext.New(ctx, codes.FailedPrecondition).AddDetailF("bot stopped before")
	}
	running = false
	return nil
}

func IsRunning() bool {
	return running
}

func handleUpdate(update tgbotapi.Update) {
	ctx, fn := context.WithTimeout(context.Background(), time.Minute)
	defer func() {
		fn()
		if r := recover(); r != nil {
			log.WithError(errors.New("an error thrown on handle update")).Error(r)
		}
	}()
	var chatID int64
	var messageID int
	if update.ChannelPost != nil {
		err = h.HandleChannelPost(ctx, update.ChannelPost)
		chatID = update.ChannelPost.Chat.ID
		messageID = update.ChannelPost.MessageID
	} else if update.CallbackQuery != nil {
		err = h.HandleCallbackQuery(ctx, update.CallbackQuery)
		chatID = update.CallbackQuery.Message.Chat.ID
		messageID = update.CallbackQuery.Message.MessageID
	} else if update.Message != nil {
		if update.Message.IsCommand() {
			err = h.HandleCommand(ctx, update.Message)
		} else {
			err = h.HandleMessage(ctx, update.Message)
		}
		chatID = update.Message.Chat.ID
		messageID = update.Message.MessageID
	}
	if err != nil {
		log.WithError(err).Error("failed to handle update")
		if e, ok := err.(*errorext.Error); ok {
			textMessage := NewTextMessage(chatID, messageID, e.Message(), nil)
			err = SendMessage(ctx, textMessage)
			if err != nil {
				log.WithError(err).Error("failed to send error to client")
			}
		} else {
			log.WithError(err).Error("message handling error")
		}
	}
}

//func SendMessage(ctx context.Context, chatId int64, message *Chattable, replyId int, keyboard *Keyboard) (tgbotapi.Message, error) {
//	if !running {
//		return nil, errorext.New(ctx, codes.Aborted)
//	}
//	switch message.Type {
//	case TextMessage:
//		log.Info("send text message")
//		msg := tgbotapi.NewMessage(chatId, message.Text)
//		if replyId > 0 {
//			msg.ReplyToMessageID = replyId
//		}
//		if keyboard != nil {
//			if keyboard.UseInline {
//				msg.ReplyMarkup = keyboard.Inline
//			} else if keyboard.UseKeyboard {
//				msg.ReplyMarkup = keyboard.Keyboard
//			}
//		}
//		r, _ := regexp.Compile("[[0-9]*]\\(tg:\\/\\/user\\?id=[0-9]*\\)")
//		if r.MatchString(message.Text) {
//			msg.ParseMode = "Markdown"
//		}
//		return bot.Send(msg)
//	case VideoMessage:
//		msg := tgbotapi.NewVideoShare(chatId, message.FileId)
//		if replyId > 0 {
//			msg.ReplyToMessageID = replyId
//		}
//		if keyboard != nil {
//			if keyboard.UseInline {
//				msg.ReplyMarkup = keyboard.Inline
//			} else if keyboard.UseKeyboard {
//				msg.ReplyMarkup = keyboard.Keyboard
//			}
//		}
//		return bot.Send(msg)
//	case AudioMessage:
//		msg := tgbotapi.NewAudioShare(chatId, message.FileId)
//		if replyId > 0 {
//			msg.ReplyToMessageID = replyId
//		}
//		if keyboard != nil {
//			if keyboard.UseInline {
//				msg.ReplyMarkup = keyboard.Inline
//			} else if keyboard.UseKeyboard {
//				msg.ReplyMarkup = keyboard.Keyboard
//			}
//		}
//		return bot.Send(msg)
//	case VoiceMessage:
//		msg := tgbotapi.NewVoiceShare(chatId, message.FileId)
//		if replyId > 0 {
//			msg.ReplyToMessageID = replyId
//		}
//		if keyboard != nil {
//			if keyboard.UseInline {
//				msg.ReplyMarkup = keyboard.Inline
//			} else if keyboard.UseKeyboard {
//				msg.ReplyMarkup = keyboard.Keyboard
//			}
//		}
//		return bot.Send(msg)
//	case PhotoMessage:
//		msg := tgbotapi.NewPhotoShare(chatId, message.FileId)
//		if replyId > 0 {
//			msg.ReplyToMessageID = replyId
//		}
//		if keyboard != nil {
//			if keyboard.UseInline {
//				msg.ReplyMarkup = keyboard.Inline
//			} else if keyboard.UseKeyboard {
//				msg.ReplyMarkup = keyboard.Keyboard
//			}
//		}
//		return bot.Send(msg)
//	case ContactMessage:
//		var c struct {
//			FirstName   string
//			PhoneNumber string
//		}
//		err := json.Unmarshal([]byte(message.MetaData), &c)
//		if err != nil {
//			return tgbotapi.Message{}, err
//		}
//		msg := tgbotapi.NewContact(chatId, c.PhoneNumber, c.FirstName)
//		if replyId > 0 {
//			msg.ReplyToMessageID = replyId
//		}
//		if keyboard != nil {
//			if keyboard.UseInline {
//				msg.ReplyMarkup = keyboard.Inline
//			} else if keyboard.UseKeyboard {
//				msg.ReplyMarkup = keyboard.Keyboard
//			}
//		}
//		return bot.Send(msg)
//	case StickerMessage:
//		msg := tgbotapi.NewStickerShare(chatId, message.FileId)
//		if replyId > 0 {
//			msg.ReplyToMessageID = replyId
//		}
//		if keyboard != nil {
//			if keyboard.UseInline {
//				msg.ReplyMarkup = keyboard.Inline
//			} else if keyboard.UseKeyboard {
//				msg.ReplyMarkup = keyboard.Keyboard
//			}
//		}
//		return bot.Send(msg)
//	case UnknownMessage:
//		msg := tgbotapi.NewMessage(chatId, "Unknown message")
//		if replyId > 0 {
//			msg.ReplyToMessageID = replyId
//		}
//		if keyboard != nil {
//			if keyboard.UseInline {
//				msg.ReplyMarkup = keyboard.Inline
//			} else if keyboard.UseKeyboard {
//				msg.ReplyMarkup = keyboard.Keyboard
//			}
//		}
//		return bot.Send(msg)
//	case LocationMessage:
//		var c struct {
//			Latitude  float64
//			Longitude float64
//		}
//		err := json.Unmarshal([]byte(message.MetaData), &c)
//		if err != nil {
//			return tgbotapi.Message{}, err
//		}
//		msg := tgbotapi.NewLocation(chatId, c.Latitude, c.Longitude)
//		if replyId > 0 {
//			msg.ReplyToMessageID = replyId
//		}
//		if keyboard != nil {
//			if keyboard.UseInline {
//				msg.ReplyMarkup = keyboard.Inline
//			} else if keyboard.UseKeyboard {
//				msg.ReplyMarkup = keyboard.Keyboard
//			}
//		}
//		return bot.Send(msg)
//	case DocumentMessage:
//		msg := tgbotapi.NewDocumentShare(chatId, message.FileId)
//		msg.Caption = message.Text
//		if replyId > 0 {
//			msg.ReplyToMessageID = replyId
//		}
//		if keyboard != nil {
//			if keyboard.UseInline {
//				msg.ReplyMarkup = keyboard.Inline
//			} else if keyboard.UseKeyboard {
//				msg.ReplyMarkup = keyboard.Keyboard
//			}
//		}
//		return bot.Send(msg)
//	case VideoNoteMessage:
//		var length int
//		err := json.Unmarshal([]byte(message.MetaData), &length)
//		if err != nil {
//			return tgbotapi.Message{}, err
//		}
//		msg := tgbotapi.NewVideoNoteShare(chatId, length, message.FileId)
//		if replyId > 0 {
//			msg.ReplyToMessageID = replyId
//		}
//		if keyboard != nil {
//			if keyboard.UseInline {
//				msg.ReplyMarkup = keyboard.Inline
//			} else if keyboard.UseKeyboard {
//				msg.ReplyMarkup = keyboard.Keyboard
//			}
//		}
//		return bot.Send(msg)
//	}
//	return &tgbotapi.Message{}, nil
//}
