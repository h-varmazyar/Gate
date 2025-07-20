package sahamyabStream

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/h-varmazyar/Gate/services/gather/configs"
	"github.com/h-varmazyar/Gate/services/gather/internal/brokers/producer"
	"github.com/h-varmazyar/Gate/services/gather/internal/models"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

type PostRepository interface {
	Create(ctx context.Context, post models.Post) error
}

type SahamyabStream struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	logger     *log.Logger
	cfg        configs.WorkerSahamyabStream
	producer   *producer.Producer
	postRepo   PostRepository

	Running bool
}

func NewWorker(
	logger *log.Logger,
	configs configs.WorkerSahamyabStream,
	producer *producer.Producer,
	postRepo PostRepository,
) *SahamyabStream {

	return &SahamyabStream{
		logger:   logger,
		cfg:      configs,
		producer: producer,
		postRepo: postRepo,
	}
}

func (w *SahamyabStream) Start() error {
	if w.cfg.Running {
		w.logger.Infof("starting tweet reader")
		w.createContext()

		if err := w.runChromeDP(); err != nil {
			return err
		}

		w.registerTweetReader()
	}

	return nil
}

func (w *SahamyabStream) Stop() {
	if w.Running {
		w.logger.Infof("stopping tweet reader")
		w.cancelFunc()
		w.Running = false
	}
}

func (w *SahamyabStream) createContext() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // نمایش مرورگر
		chromedp.Flag("disable-gpu", true),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	w.ctx, w.cancelFunc = chromedp.NewContext(allocCtx)

	//w.ctx, w.cancelFunc = chromedp.NewContext(context.Background())
}

func (w *SahamyabStream) runChromeDP() error {
	if err := chromedp.Run(w.ctx, network.Enable()); err != nil {
		return err
	}
	return nil
}

func (w *SahamyabStream) registerTweetReader() {
	var targetWSID string
	var targetWSURL = "wss://push.sahamyab.com/lightstreamer"

	chromedp.ListenTarget(w.ctx, func(ev interface{}) {
		switch e := ev.(type) {
		case *network.EventWebSocketCreated:
			fmt.Println("🧩 WebSocket Created:", e.URL)
			if strings.HasPrefix(e.URL, targetWSURL) {
				targetWSID = e.RequestID.String()
				fmt.Println("🎯 Found target WebSocket:", targetWSID)
			}

		case *network.EventWebSocketFrameReceived:
			fmt.Println("<UNK> WebSocket FrameReceived:", e.RequestID)
			if e.RequestID.String() != targetWSID {
				return
			}

			fmt.Println(e.Response.PayloadData)
			items := responseParser(e.Response.PayloadData)
			fmt.Println(items)
			if len(items) < 9 {
				return
			}

			fmt.Println("new item")

			if err := w.preparePost(items); err != nil {
				w.logger.Error(err)
			}
		}
	})

	err := chromedp.Run(w.ctx, chromedp.Navigate("https://www.sahamyab.com/stocktwits"))
	if err != nil {
		log.Fatal(err)
	}
}

func (w *SahamyabStream) preparePost(items []string) error {
	post := models.Post{
		Provider: "SAHAMYAB",
	}
	id, err := strconv.Atoi(items[2])
	if err != nil {
		return err
	}

	post.ID = uint(id)

	switch items[3] {
	case "9":
		post.SenderUsername = items[5]
		post.Content = items[6]
		post.PostedAt, _ = time.Parse(time.RFC3339, items[7])
	case "11":
		post.Content = items[4]
		post.PostedAt, _ = time.Parse(time.RFC3339, items[5])
	}

	post.Tags = fetchTags(post.Content)

	if err := w.postRepo.Create(w.ctx, post); err != nil {
		w.logger.WithError(err).Error("failed to save stream post")
		return err
	}

	postPayload := producer.PostPayload{
		PostedAt:       post.PostedAt,
		Id:             post.ID,
		Content:        post.Content,
		SenderUsername: post.SenderUsername,
		Provider:       string(post.Provider),
		Tags:           post.Tags,
	}
	if post.LikeCount != nil {
		postPayload.LikeCount = *post.LikeCount
	}
	if err := w.producer.PublishPost(postPayload); err != nil {
		w.logger.WithError(err).Error("failed to produce ticker")
		return err
	}
	return nil
}

func responseParser(rawData string) []string {
	data := make([]string, 0)
	items := strings.Split(rawData, ");")
	for _, item := range items {
		item = strings.TrimSpace(item)
		if strings.HasPrefix(item, "d(") {
			rawData = item
			break
		}
	}
	newData := ""
	inString := false
	for i := 2; i < len(rawData); i++ {
		if rawData[i] == ',' {
			if inString {
				newData += string(rune(rawData[i]))
			} else {
				data = append(data, decodeUnicode(newData))
				newData = ""
			}
		} else if rawData[i] == '\'' {
			inString = !inString
		} else {
			newData += string(rune(rawData[i]))
		}
	}

	data = append(data, decodeUnicode(newData))

	return data
}

func decodeUnicode(s string) string {
	quoted := fmt.Sprintf(`"%s"`, s)
	var decoded string
	if err := json.Unmarshal([]byte(quoted), &decoded); err != nil {
		fmt.Println(err)
		return s
	}
	return decoded
}

func fetchTags(content string) []string {
	tags := make([]string, 0)
	for _, word := range strings.Split(content, " ") {
		if strings.HasPrefix(word, "#") {
			tags = append(tags, word)
		}
	}
	return tags
}
