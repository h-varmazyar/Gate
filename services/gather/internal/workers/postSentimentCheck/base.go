package postSentimentCheck

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/h-varmazyar/Gate/services/gather/configs"
	"github.com/h-varmazyar/Gate/services/gather/internal/models"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/genai"
	"io"
	"net/http"
	"strings"
	"time"
)

type PostRepository interface {
	List(ctx context.Context) ([]models.Post, error)
	UpdateSentiment(ctx context.Context, postID uint, sentiment float64) error
}

var (
	sentimentDetectionPrompt = `
please select the polarity of below tweets between -1 to 1. 
tweets prepared based on json model with its id(tid) and text.
the response based on below json:
{
polarities:[
{
"tid":{tid_value},
"p":{polarity_value}
},
]
}
%v
`
)

var geminiModels = []string{
	"gemini-2.5-pro",
	"gemini-2.5-flash",
	//"gemini-2.5-flash-light",
	"gemini-2.0-flash",
	//"gemini-2.0-flash-light",
}

type PostSentimentCheck struct {
	ctx          context.Context
	cancelFunc   context.CancelFunc
	logger       *log.Logger
	cfg          configs.WorkerPostSentimentCheck
	postRepo     PostRepository
	geminiClient *genai.Client
	client       *http.Client
	modelIndex   int

	Running bool
}

type tweet struct {
	TweetID  uint    `json:"tid"`
	Text     string  `json:"text"`
	Polarity float64 `json:"polarity,omitempty"`
}

type tweets struct {
	Tweets []tweet `json:"tweets"`
}

func NewWorker(
	logger *log.Logger,
	cfg configs.WorkerPostSentimentCheck,
	postRepo PostRepository,
) (*PostSentimentCheck, error) {

	w := &PostSentimentCheck{
		logger:   logger,
		cfg:      cfg,
		postRepo: postRepo,
		client:   &http.Client{},
	}

	w.ctx, w.cancelFunc = context.WithCancel(context.Background())

	var err error
	//dialer, err := proxy.SOCKS5("tcp", cfg.SocksProxyAddress, nil, proxy.Direct)
	//if err != nil {
	//	return nil, err
	//}

	httpTransport := &http.Transport{
		TLSHandshakeTimeout: 50 * time.Second,
		//DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
		//	return dialer.Dial(network, addr)
		//},
	}

	client := &http.Client{
		Transport: httpTransport,
		Timeout:   50 * time.Second,
	}

	w.geminiClient, err = genai.NewClient(w.ctx, &genai.ClientConfig{
		APIKey:     cfg.GeminiAPIKey,
		Backend:    genai.BackendGeminiAPI,
		HTTPClient: client,
	})
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (w *PostSentimentCheck) Start() error {
	if w.cfg.Running {
		w.logger.Infof("starting archive tweet reader")
		go w.run()
	}
	return nil
}

func (w *PostSentimentCheck) Stop() {
	if w.Running {
		w.logger.Infof("stopping archive tweet reader")
		w.cancelFunc()
		w.Running = false
	}
}

func (w *PostSentimentCheck) run() {
	ticker := time.NewTicker(time.Second * 50)

	for {
		select {
		case <-w.ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			fmt.Println("checking tweets polarity")
			posts, err := w.postRepo.List(w.ctx)
			if err != nil {
				w.logger.WithError(err).Errorf("failed to list posts")
				continue
			}

			polarityMap, err := w.detectPolarity(posts)
			if err != nil {
				w.logger.WithError(err).Error("failed to detect polarity")
				continue
			}

			for _, polarity := range polarityMap.Polarities {
				//id, err := strconv.ParseUint(polarity.Tid, 10, 64)
				//if err != nil {
				//	continue
				//}
				if err = w.postRepo.UpdateSentiment(w.ctx, uint(polarity.Tid), polarity.Polarity); err != nil {
					w.logger.WithError(err).Errorf("failed to update sentiment")
				}
			}
		}
	}
}

type geminiBody struct {
	Contents []content `json:"contents"`
}

type geminiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
	Details []struct {
		Type       string `json:"@type"`
		Violations []struct {
			QuotaMetric     string `json:"quotaMetric"`
			QuotaID         string `json:"quotaId"`
			QuotaDimensions struct {
				Location string `json:"location"`
				Model    string `json:"model"`
			} `json:"quotaDimensions"`
			QuotaValue string `json:"quotaValue"`
		} `json:"violations,omitempty"`
		Links []struct {
			Description string `json:"description"`
			URL         string `json:"url"`
		} `json:"links,omitempty"`
		RetryDelay string `json:"retryDelay,omitempty"`
	} `json:"details"`
}

type geminiResp struct {
	Error      *geminiError `json:"error"`
	Candidates []candidate  `json:"candidates"`
}

type candidate struct {
	Content content `json:"content"`
}
type content struct {
	Parts []part `json:"parts"`
}

type part struct {
	Text string `json:"text"`
}

type tweetResp struct {
	Polarities []struct {
		Tid      int     `json:"tid"`
		Polarity float64 `json:"p"`
	} `json:"polarities"`
}

func (w *PostSentimentCheck) detectPolarity(posts []models.Post) (*tweetResp, error) {
	endpoint := fmt.Sprintf("https://generativelanguage.googleapis.com//v1beta/models/%v:generateContent?key=%v", geminiModels[w.modelIndex], w.cfg.GeminiAPIKey)

	tweets := tweets{
		Tweets: make([]tweet, 0),
	}
	for _, post := range posts {
		tweets.Tweets = append(tweets.Tweets, tweet{
			TweetID: post.ID,
			Text:    post.Content,
		})
	}

	tweetsJSON, err := json.Marshal(tweets)
	if err != nil {
		return nil, err
	}

	body := geminiBody{
		Contents: []content{
			{
				Parts: []part{
					{
						Text: fmt.Sprintf(sentimentDetectionPrompt, string(tweetsJSON)),
					},
				},
			},
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	resp, err := w.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// fmt.Println(string(respBody))

	gr := new(geminiResp)

	err = json.Unmarshal(respBody, gr)
	if err != nil {
		return nil, err
	}

	if gr.Error != nil {
		if gr.Error.Code == 429 {
			w.nextQuota()
		}
		return nil, fmt.Errorf(gr.Error.Message)
	}
	text := gr.Candidates[0].Content.Parts[0].Text

	text = strings.TrimPrefix(text, "```json")
	text = strings.TrimSuffix(text, "```")

	respMap := new(tweetResp)
	err = json.Unmarshal([]byte(text), &respMap)
	if err != nil {
		return nil, err
	}

	return respMap, nil
}

func (w *PostSentimentCheck) nextQuota() {
	w.modelIndex++
	if w.modelIndex == len(geminiModels) {
		w.modelIndex = 0
	}
}

//func (w *PostSentimentCheck) detectPolarity(posts []models.Post) (map[string]float64, error) {
//	w.logger.Infof("detecting polarity of tweets")
//
//	//chat, err := w.geminiClient.Chats.Create(w.ctx, "gemini-2.5-flash", new(genai.GenerateContentConfig), nil)
//	//if err != nil {
//	//	return nil, err
//	//}
//
//	tweets := tweets{
//		Tweets: make([]tweet, 0),
//	}
//	for _, post := range posts {
//		tweets.Tweets = append(tweets.Tweets, tweet{
//			TweetID: post.ID,
//			Text:    post.Content,
//		})
//	}
//
//	tweetsJSON, err := json.Marshal(tweets)
//	if err != nil {
//		return nil, err
//	}
//
//	fmt.Println(fmt.Sprintf(sentimentDetectionPrompt, string(tweetsJSON)))
//
//	contents := []*genai.Content{
//		{
//			Parts: []*genai.Part{
//				{
//					Text: fmt.Sprintf(sentimentDetectionPrompt, string(tweetsJSON)),
//				},
//			},
//		},
//	}
//
//	contentResp, err := w.geminiClient.Models.GenerateContent(w.ctx, "gemini-2.5-flash", contents, new(genai.GenerateContentConfig))
//	if err != nil {
//		return nil, err
//	}
//
//	textResp := contentResp.Text()
//
//	respMap := make(map[string]float64)
//
//	err = json.Unmarshal([]byte(textResp), &respMap)
//	if err != nil {
//		return nil, err
//	}
//
//	return respMap, nil
//}
