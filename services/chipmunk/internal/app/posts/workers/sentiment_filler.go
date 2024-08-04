package workers

import (
	"errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/llmAssists"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

var (
	ErrDetectorLimitReach = errors.New("detector limit reach")
)

type SentimentDetector interface {
	DetectEmotions(ctx context.Context, posts []*entity.Post) ([]*entity.Post, error)
}

type SentimentFillerRepository interface {
	SavePost(ctx context.Context, post *entity.Post) error
	ReturnEmotionlessPost(ctx context.Context, maxLength int64) (*entity.Post, error)
}

type SentimentFiller struct {
	configs           Configs
	log               *log.Logger
	sentimentDetector SentimentDetector
	repository        SentimentFillerRepository
}

func NewSentimentFiller(log *log.Logger, configs Configs, repository SentimentFillerRepository) *SentimentFiller {
	return &SentimentFiller{
		configs:    configs,
		log:        log,
		repository: repository,
	}
}

func (w *SentimentFiller) Start(ctx context.Context) {
	w.log.Infof("starting sentiment filler worker")
	networkConn := grpcext.NewConnection(w.configs.NetworkAddress)
	networkService := networkAPI.NewRequestServiceClient(networkConn)
	w.sentimentDetector = llmAssists.NewGemini(w.log, w.configs.SentimentDetectorToken, networkService)

	go w.detectPostsSentiment(ctx)
}

func (w *SentimentFiller) detectPostsSentiment(ctx context.Context) {
	w.log.Infof("detecting emotionless posts sentiment")

	ticker := time.NewTicker(time.Second * 2)
	for {
		select {
		case <-ticker.C:
			emotionlessPosts := make([]*entity.Post, 0)
			remainingCharLength := w.configs.MaxSentimentDetectorTokenLength
			for remainingCharLength > 0 {
				emotionlessPost, err := w.repository.ReturnEmotionlessPost(ctx, remainingCharLength)
				if err != nil {
					break
				}
				remainingCharLength -= int64(len(emotionlessPost.Content))
			}

			emotionalPosts, err := w.sentimentDetector.DetectEmotions(ctx, emotionlessPosts)
			if err != nil {
				w.log.WithError(err).Error("failed to detect emotion of post")
				w.handleDetectorError(err)
				continue
			}

			for _, post := range emotionalPosts {
				if err = w.repository.SavePost(ctx, post); err != nil {
					w.log.WithError(err).Errorf("failed to save emotional post: %v", post.Id)
				}
			}
		}
	}
}

func (w *SentimentFiller) handleDetectorError(err error) {
	switch {
	case errors.Is(err, ErrDetectorLimitReach):
		time.Sleep(time.Minute * 2)
	}
}
