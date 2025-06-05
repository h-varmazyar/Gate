package tweetReader

import (
	"encoding/json"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/h-varmazyar/Gate/services/gather/configs"
	"github.com/h-varmazyar/Gate/services/gather/internal/brokers/producer"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"strings"
	"time"
)

type TweetReader struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	logger     *log.Logger
	cfg        configs.WorkerTweetReader
	producer   *producer.Producer

	Running bool
}

func NewWorker(
	logger *log.Logger,
	configs configs.WorkerTweetReader,
	producer *producer.Producer,
) *TweetReader {

	return &TweetReader{
		logger:   logger,
		cfg:      configs,
		producer: producer,
	}
}

const (
//rawData = `d(7,1,'458016341',11,'content','2025-06-04T13:06:54Z',14,'1404/03/14 16:36',16);s(7,1);`
//rawData = "d(7,1,'458016359',9,'mahdi','mahdi8145','content','2025-06-04T13:08:15Z',4,'pull','bf4a733c-635b-47e4-b84d-49e469250938',8,'1404/03/14 16:38',3,'[{\"id\":\"458016364\",\"optionBody\":\"\\u0628\\u0644\\u0647\",\"score\":0.0,\"userChecked\":false},{\"id\":\"458016365\",\"optionBody\":\"\\u062E\\u06CC\\u0631\",\"score\":0.0,\"userChecked\":false}]','1404/03/22 15:38','Jun 12, 2025 3:38:15 PM','22.0','7.0',4,'continue',3);s(7,1);"
//rawData = "d(7,1,'458016380',9,'content','2025-06-04T13:10:56Z',4,'twit','default',8,'1404/03/14 16:40',3,'#','$','#','#','#',4,'#',3);"
//rawData = "d(7,1,'458016414',11,'content','2025-06-04T13:16:41Z',14,'1404/03/14 16:46',16);s(7,1);"
//rawData = "d(7,1,'458016392',11,'content','2025-06-04T13:12:13Z',14,'1404/03/14 16:42',16);s(7,1);"

// rawData = `d(7,1,'458016341',11,'##\\u0641\\u0627\\u0631\\u0627\\u06A9 \\u0645\\u062A\\u0627\\u0633\\u0641\\u0627\\u0646\\u0647 \\u0645\\u0634\\u06A9\\u0644 \\u0627\\u0632 \\u0634\\u0646\\u0627\\u0648\\u0631 \\u0632\\u06CC\\u0627\\u062F \\u0648 \\u062D\\u0642\\u06CC\\u0642\\u06CC \\u0647\\u0627\\u06CC \\u062A\\u0631\\u0633\\u0648 \\u0647\\u0633\\u062A \\u06A9\\u0647 \\u0641\\u0642\\u0637 \\u0628\\u0644\\u062F\\u0646 \\u0628\\u0631\\u0646 \\u062A\\u0648 \\u0635\\u0641 \\u0641\\u0631\\u0648\\u0634','2025-06-04T13:06:54Z',14,'1404/03/14 16:36',16);s(7,1);`
// rawData = "d(7,1,'458016359',9,'mahdi','mahdi8145','##\\u0646\\u0638\\u0631\\u0633\\u0646\\u062C\\u06CC \\u000A\\u000A\\u0622\\u06CC\\u0627 \\u062A\\u0627\\u06A9\\u0646\\u0648\\u0646 \\u0633\\u06CC\\u0628 \\u0648 \\u06AF\\u0644\\u0627.\\u0628\\u06CC \\u0628\\u0631.\\u062C\\u0627\\u0645 (\\u0646\\u0645\\u0648\\u0646\\u0647 \\u0627\\u0632 \\u0645\\u0630\\u0627\\u06A9\\u0631\\u0627\\u062A\\u0650 \\u0645\\u0648\\u0641\\u0642!) \\u0631\\u0627 \\u062E\\u0648\\u0631\\u062F\\u0647 \\u0627\\u06CC\\u062F\\u061F\\u000A\\u000A #\\u0634\\u0627\\u062E\\u0635_\\u0628\\u0648\\u0631\\u0633 #\\u0639\\u06CC\\u0627\\u0631','2025-06-04T13:08:15Z',4,'pull','bf4a733c-635b-47e4-b84d-49e469250938',8,'1404/03/14 16:38',3,'[{\"id\":\"458016364\",\"optionBody\":\"\\u0628\\u0644\\u0647\",\"score\":0.0,\"userChecked\":false},{\"id\":\"458016365\",\"optionBody\":\"\\u062E\\u06CC\\u0631\",\"score\":0.0,\"userChecked\":false}]','1404/03/22 15:38','Jun 12, 2025 3:38:15 PM','22.0','7.0',4,'continue',3);s(7,1);"
// rawData = "d(7,1,'458016380',9,'\\u062A\\u0627\\u067E\\u06CC\\u06A9\\u0648','reza1477','##\\u0641\\u0648\\u0644\\u0627\\u062F\\u000A\\u2666\\uFE0F\\u0627\\u062F\\u0639\\u0627\\u06CC \\u0622\\u06A9\\u0633\\u06CC\\u0648\\u0633 \\u062F\\u0631\\u0628\\u0627\\u0631\\u0647 \\u0646\\u0638\\u0631 \\u0627\\u06CC\\u0631\\u0627\\u0646 \\u062D\\u0648\\u0644 \\u06A9\\u0646\\u0633\\u0631\\u0633\\u06CC\\u0648\\u0645 \\u063A\\u0646\\u06CC\\u200C\\u0633\\u0627\\u0632\\u06CC\\u000A\\u000A\\uD83D\\uDD39\\u0648\\u0628\\u06AF\\u0627\\u0647 \\u062E\\u0628\\u0631\\u06CC-\\u062A\\u062D\\u0644\\u06CC\\u0644\\u06CC \\u00AB\\u0622\\u06A9\\u0633\\u06CC\\u0648\\u0633\\u00BB \\u0628\\u0627\\u0645\\u062F\\u0627\\u062F \\u0686\\u0647\\u0627\\u0631\\u0634\\u0646\\u0628\\u0647 \\u0628\\u0647 \\u0646\\u0642\\u0644 \\u0627\\u0632 \\u0645\\u0642\\u0627\\u0645\\u200C\\u0647\\u0627\\u06CC \\u0627\\u06CC\\u0631\\u0627\\u0646\\u06CC \\u0645\\u062F\\u0639\\u06CC \\u0634\\u062F \\u06A9\\u0647 \\u062A\\u0647\\u0631\\u0627\\u0646 \\u0634\\u0627\\u06CC\\u062F \\u0628\\u0627 \\u067E\\u06CC\\u0634\\u0646\\u0647\\u0627\\u062F \\u0648\\u0627\\u0634\\u0646\\u06AF\\u062A\\u0646 \\u062F\\u0631\\u0628\\u0627\\u0631\\u0647 \\u0627\\u06CC\\u062C\\u0627\\u062F \\u06A9\\u0646\\u0633\\u0631\\u0633\\u06CC\\u0648\\u0645 \\u063A\\u0646\\u06CC\\u200C\\u0633\\u0627\\u0632\\u06CC \\u0645\\u0646\\u0637\\u0642\\u0647\\u200C\\u0627\\u06CC \\u0645\\u0648\\u0627\\u0641\\u0642\\u062A \\u06A9\\u0646\\u062F\\u060C \\u062F\\u0631 \\u0635\\u0648\\u0631\\u062A\\u06CC \\u06A9\\u0647 \\u0627\\u06CC\\u0646 \\u06A9\\u0646\\u0633\\u0631\\u0633\\u06CC\\u0648\\u0645 \\u062F\\u0627\\u062E\\u0644 \\u0627\\u06CC\\u0631\\u0627\\u0646 \\u0627\\u06CC\\u062C\\u0627\\u062F \\u0634\\u0648\\u062F.\\u000A\\uD83D\\uDD39\\u0627\\u06CC\\u0646 \\u0648\\u0628\\u06AF\\u0627\\u0647 \\u0627\\u062F\\u0639\\u0627 \\u06A9\\u0631\\u062F \\u06A9\\u0647 \\u0627\\u06CC\\u0646 \\u067E\\u0627\\u0633\\u062E \\u0646\\u0634\\u0627\\u0646 \\u0645\\u06CC\\u200C\\u062F\\u0647\\u062F \\u06A9\\u0647 \\u062A\\u0647\\u0631\\u0627\\u0646 \\u0645\\u0645\\u06A9\\u0646 \\u0627\\u0633\\u062A \\u067E\\u06CC\\u0634\\u0646\\u0647\\u0627\\u062F \\u0648\\u06CC\\u062A\\u06A9\\u0627\\u0641 \\u0631\\u0627 \\u0628\\u0647\\u200C\\u0637\\u0648\\u0631 \\u06A9\\u0627\\u0645\\u0644 \\u0631\\u062F \\u0646\\u06A9\\u0646\\u062F\\u060C \\u0628\\u0644\\u06A9\\u0647 \\u062F\\u0631 \\u0639\\u0648\\u0636 \\u0628\\u0647 \\u062F\\u0646\\u0628\\u0627\\u0644 \\u0645\\u0630\\u0627\\u06A9\\u0631\\u0647 \\u062F\\u0631\\u0645\\u0648\\u0631\\u062F \\u062C\\u0632\\u0626\\u06CC\\u0627\\u062A \\u0628\\u0627\\u0634\\u062F.\\u000A\\uD83D\\uDD39\\u0622\\u06A9\\u0633\\u06CC\\u0648\\u0633 \\u0647\\u0645\\u0686\\u0646\\u06CC\\u0646 \\u0628\\u0647 \\u0646\\u0642\\u0644 \\u0627\\u0632 \\u0645\\u0646\\u0627\\u0628\\u0639\\u06CC \\u0645\\u062F\\u0639\\u06CC \\u0634\\u062F \\u06A9\\u0647 \\u062F\\u0648\\u0631 \\u0634\\u0634\\u0645 \\u0645\\u0630\\u0627\\u06A9\\u0631\\u0627\\u062A \\u063A\\u06CC\\u0631\\u0645\\u0633\\u062A\\u0642\\u06CC\\u0645 \\u0628\\u06CC\\u0646 \\u0627\\u06CC\\u0631\\u0627\\u0646 \\u0648 \\u0627\\u0645\\u0631\\u06CC\\u06A9\\u0627 \\u0645\\u0645\\u06A9\\u0646 \\u0627\\u0633\\u062A \\u062F\\u0631 \\u0622\\u062E\\u0631 \\u0627\\u06CC\\u0646 \\u0647\\u0641\\u062A\\u0647 \\u062F\\u0631 \\u0645\\u06A9\\u0627\\u0646\\u06CC \\u062F\\u0631 \\u063A\\u0631\\u0628 \\u0622\\u0633\\u06CC\\u0627 \\u0628\\u0631\\u06AF\\u0632\\u0627\\u0631 \\u0634\\u0648\\u062F.\\u000A@AkhbareFori \\u007C Link','2025-06-04T13:10:56Z',4,'twit','default',8,'1404/03/14 16:40',3,'#','$','#','#','#',4,'#',3);"
// rawData = "d(7,1,'458016414',11,'##\\u0634\\u062A\\u0631\\u0627\\u0646\\u000A\\u0648\\u0627\\u0634\\u06CC\\u0646\\u06AF\\u062A\\u0646 \\u067E\\u0633\\u062A : \\u000A\\u0631\\u0647\\u0628 \\u0631 \\u0627\\u06CC\\u0631\\u0627\\u0646 \\u0628\\u0647 \\u0622\\u0645\\u0631\\u06CC\\u06A9\\u0627 \\u0627\\u06CC\\u0631\\u0627\\u062F \\u06AF\\u0631\\u0641\\u062A \\u0648\\u0644\\u06CC \\u0627\\u06CC\\u062F\\u0647 \\u0647\\u0627\\u06CC \\u062A\\u0648\\u0627\\u0641\\u0642 \\u0631\\u0627 \\u0631\\u062F \\u0646\\u06A9\\u0631\\u062F\\u0647 \\u0627\\u0633\\u062A','2025-06-04T13:16:41Z',14,'1404/03/14 16:46',16);s(7,1);"
// rawData = "d(7,1,'458016392',11,'## \\u0641\\u0646 \\u0627\\u0641\\u0632\\u0627\\u0631\\u000A\\u0641\\u0627\\u06CC\\u0646\\u0646\\u0634\\u0627\\u0644 \\u062A\\u0627\\u06CC\\u0645\\u0632\\u060C \\u0628\\u0647 \\u0646\\u0642\\u0644 \\u0627\\u0632 \\u06CC\\u06A9 \\u0645\\u0642\\u0627\\u0645 \\u0627\\u0631\\u0648\\u067E\\u0627\\u06CC\\u06CC: \\u000A\\u062F\\u0648\\u0644\\u062A \\u062A\\u0631\\u0627\\u0645\\u067E \\u0628\\u0647 \\u0627\\u064A\\u0631\\u0627\\u0646 \\u0627\\u0637\\u0644\\u0627\\u0639 \\u062F\\u0627\\u062F \\u06A9\\u0647 \\u0637\\u0628\\u0642 \\u062A\\u0648\\u0627\\u0641\\u0642 \\u0645\\u0648\\u0642\\u062A\\u060C \\u0628\\u0647 \\u0627\\u06CC\\u0646 \\u06A9\\u0634\\u0648\\u0631 \\u0627\\u062C\\u0627\\u0632\\u0647 \\u063A\\u0646\\u06CC\\u200C\\u0633\\u0627\\u0632\\u06CC \\u062F\\u0631 \\u0633\\u0637\\u062D \\u067E\\u0627\\u06CC\\u06CC\\u0646 \\u0631\\u0627 \\u0645\\u06CC\\u200C\\u062F\\u0647\\u062F.','2025-06-04T13:12:13Z',14,'1404/03/14 16:42',16);s(7,1);"
)

func (w *TweetReader) Start() error {
	//fmt.Println(responseParser(rawData))
	//return nil

	w.logger.Infof("starting tweet reader")
	w.createContext()

	if err := w.runChromeDP(); err != nil {
		return err
	}

	w.registerTweetReader()
	return nil
}

func (w *TweetReader) Stop() {
	if w.Running {
		w.logger.Infof("stopping tweet reader")
		w.cancelFunc()
		w.Running = false
	}
}

func (w *TweetReader) createContext() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // نمایش مرورگر
		chromedp.Flag("disable-gpu", true),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	w.ctx, w.cancelFunc = chromedp.NewContext(allocCtx)

	//w.ctx, w.cancelFunc = chromedp.NewContext(context.Background())
}

func (w *TweetReader) runChromeDP() error {
	if err := chromedp.Run(w.ctx, network.Enable()); err != nil {
		return err
	}
	return nil
}

func (w *TweetReader) registerTweetReader() {
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

			items := responseParser(targetWSID)
			if len(items) < 9 {
				return
			}

			postPayload := w.preparePost(items)
			if err := w.producer.PublishPost(postPayload); err != nil {
				w.logger.WithError(err).Error("failed to produce ticker")
			}
		}
	})

	err := chromedp.Run(w.ctx, chromedp.Navigate("https://www.sahamyab.com/stocktwits"))
	if err != nil {
		log.Fatal(err)
	}
}

func (w *TweetReader) preparePost(items []string) producer.PostPayload {
	post := producer.PostPayload{
		Id:       items[2],
		Provider: "SAHAMYAB",
	}

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
	return post
}

func responseParser(rawData string) []string {
	data := make([]string, 0)
	rawData = strings.Split(rawData, ");")[0]
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
