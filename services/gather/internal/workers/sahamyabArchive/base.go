package sahamyabArchive

import (
	"errors"
	"fmt"
	"github.com/h-varmazyar/Gate/services/gather/configs"
	"github.com/h-varmazyar/Gate/services/gather/internal/domain"
	"github.com/h-varmazyar/Gate/services/gather/internal/models"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"strings"
	"time"
)

type PostRepository interface {
	BatchSave(ctx context.Context, posts []*models.Post) error
	Create(ctx context.Context, post models.Post) error
	GetOldest(ctx context.Context, provider models.PostProvider) (*models.Post, error)
}

type SahamyabAdapter interface {
	GetUserPageList(ctx context.Context, input domain.GetScoredSahamyabPost) (domain.SahamyabPostList, error)
}

type SahamyabArchive struct {
	ctx             context.Context
	cancelFunc      context.CancelFunc
	logger          *log.Logger
	cfg             configs.WorkerSahamyabArchive
	postRepo        PostRepository
	sahamyabAdapter SahamyabAdapter

	Running bool
}

func NewWorker(
	logger *log.Logger,
	cfg configs.WorkerSahamyabArchive,
	postRepo PostRepository,
	sahamyabAdapter SahamyabAdapter,
) *SahamyabArchive {

	w := &SahamyabArchive{
		logger:          logger,
		cfg:             cfg,
		postRepo:        postRepo,
		sahamyabAdapter: sahamyabAdapter,
	}

	w.ctx, w.cancelFunc = context.WithCancel(context.Background())

	return w
}

func (w *SahamyabArchive) Start() error {
	if w.cfg.Running {
		w.logger.Infof("starting archive tweet reader")
		go w.run()
	}
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
	ticker := time.NewTicker(time.Minute / 2)

	page := 0
	ScoredPostDate := ""

	oldestPost, err := w.postRepo.GetOldest(w.ctx, models.PostProviderSahamyab)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			w.logger.WithError(err).Error("failed to get last posted at")
			return
		}
	} else {
		ScoredPostDate = fmt.Sprintf("%v", oldestPost.PostedAt.UnixMilli())
	}

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

			page++

		}
	}
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
