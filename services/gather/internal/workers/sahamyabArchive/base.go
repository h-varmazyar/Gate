package sahamyabArchive

import (
	"encoding/json"
	"fmt"
	"github.com/h-varmazyar/Gate/services/gather/configs"
	"github.com/h-varmazyar/Gate/services/gather/internal/domain"
	"github.com/h-varmazyar/Gate/services/gather/internal/models"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"golang.org/x/net/proxy"
	"google.golang.org/genai"
	"net"
	"net/http"
	"strings"
	"time"
)

type PostRepository interface {
	BatchSave(ctx context.Context, posts []*models.Post) error
	Create(ctx context.Context, post models.Post) error
}

type SahamyabAdapter interface {
	GetUserPageList(ctx context.Context, input domain.GetScoredSahamyabPost) (domain.SahamyabPostList, error)
}

var (
	sentimentDetectionPrompt = `
please select the polarity of below tweets based on float value between -1 to 1. 
the count of tweets are %v. tweets prepared based on json model with its id(tid) and text.
return the polarity of each tweet in the json format with the key of tid for tweet id and value with polarity.
%v
`
)

type SahamyabArchive struct {
	ctx             context.Context
	cancelFunc      context.CancelFunc
	logger          *log.Logger
	cfg             configs.WorkerSahamyabArchive
	postRepo        PostRepository
	sahamyabAdapter SahamyabAdapter
	geminiClient    *genai.Client

	Running bool
}

func NewWorker(
	logger *log.Logger,
	cfg configs.WorkerSahamyabArchive,
	postRepo PostRepository,
	sahamyabAdapter SahamyabAdapter,
) (*SahamyabArchive, error) {

	w := &SahamyabArchive{
		logger:          logger,
		cfg:             cfg,
		postRepo:        postRepo,
		sahamyabAdapter: sahamyabAdapter,
	}

	w.ctx, w.cancelFunc = context.WithCancel(context.Background())

	dialer, err := proxy.SOCKS5("tcp", cfg.SocksProxyAddress, nil, proxy.Direct)
	if err != nil {
		return nil, err
	}

	httpTransport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		},
	}

	client := &http.Client{
		Transport: httpTransport,
		Timeout:   10 * time.Second,
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

func (w *SahamyabArchive) Start() error {
	w.logger.Infof("starting archive tweet reader")
	go w.run()
	return nil
}

func (w *SahamyabArchive) Stop() {
	if w.Running {
		w.logger.Infof("stopping archive tweet reader")
		w.cancelFunc()
		w.Running = false
	}
}

func (w *SahamyabArchive) run() {
	ticker := time.NewTicker(time.Minute / 4)

	page := 0
	ScoredPostDate := ""
	for {
		select {
		case <-w.ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			scoredPosts, err := w.sahamyabAdapter.GetUserPageList(w.ctx, domain.GetScoredSahamyabPost{
				Page:           page,
				ScoredPostDate: ScoredPostDate,
			})
			if err != nil {
				w.logger.WithError(err).Error("failed to get sahamyab scored posts")
				continue
			}

			for _, item := range scoredPosts.Items {
				if !hasValidContent(item.Content) {
					continue
				}
				post := models.Post{
					PostedAt:       item.SendTime,
					Content:        item.Content,
					ParentId:       uint(item.ParentID),
					SenderUsername: item.SenderUsername,
					Provider:       models.PostProviderSahamyab,
					Tags:           fetchTags(item.Content),
					LikeCount:      &item.LikeCount,
					RetwitCount:    &item.RetwitCount,
					CommentCount:   &item.CommentCount,
					QuoteCount:     &item.QuoteCount,
					Type:           item.Type,
				}
				post.ID = uint(item.ID)

				if err := w.postRepo.Create(w.ctx, post); err != nil {
					w.logger.WithError(err).Error("failed to save post")
				}
				ScoredPostDate = item.ScoredPostDate
			}

			//posts, err = w.detectPolarity(posts)
			//if err != nil {
			//	w.logger.WithError(err).Error("failed to detect polarity")
			//}
			//
			//if err = w.postRepo.BatchSave(w.ctx, posts); err != nil {
			//	w.logger.WithError(err).Error("failed to save posts")
			//	continue
			//}

			page++

		}
	}
}

type tweet struct {
	TweetID  uint    `json:"tid"`
	Text     string  `json:"text"`
	Polarity float64 `json:"polarity"`
}

type tweets struct {
	Tweets []tweet `json:"tweets"`
}

func (w *SahamyabArchive) detectPolarity(posts []*models.Post) ([]*models.Post, error) {
	w.logger.Infof("detecting polarity of tweets")
	chat, err := w.geminiClient.Chats.Create(w.ctx, "2.5 Flash", new(genai.GenerateContentConfig), nil)
	if err != nil {
		return posts, err
	}

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
		return posts, err
	}

	contents := []*genai.Content{
		{
			Parts: []*genai.Part{
				{
					Text: fmt.Sprintf(sentimentDetectionPrompt, len(posts), tweetsJSON),
				},
			},
		},
	}
	contentResp, err := chat.GenerateContent(w.ctx, "2.5 Flash", contents, new(genai.GenerateContentConfig))
	if err != nil {
		return posts, err
	}

	textResp := contentResp.Text()

	err = json.Unmarshal([]byte(textResp), &tweets)
	if err != nil {
		return posts, err
	}

	fmt.Println(textResp)
	fmt.Println(tweets)

	return posts, nil
}

func hasValidContent(content string) bool {
	strings.ReplaceAll(content, "\n", " ")
	return len(strings.Split(content, " ")) > 2
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
