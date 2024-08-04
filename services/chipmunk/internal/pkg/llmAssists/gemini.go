package llmAssists

import (
	"errors"
	"fmt"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"strings"
)

type Gemini struct {
	log            *log.Logger
	networkService networkAPI.RequestServiceClient
	token          string
}

var errEmptyPosts = errors.New("empty posts")

var (
	requestTokenTemplate = `
  describe the sentiment of below posts in the format of json like this:
[
"id":"sentiment status(positive, negative or neutral)"
]

each posts in the format of {id}- {content}

%v
`
)

func NewGemini(log *log.Logger, token string, networkService networkAPI.RequestServiceClient) *Gemini {
	return &Gemini{
		log:            log,
		networkService: networkService,
		token:          token,
	}
}

func (g *Gemini) DetectEmotions(ctx context.Context, posts []*entity.Post) ([]*entity.Post, error) {
	if len(posts) == 0 {
		return nil, errEmptyPosts
	}

	prompt := g.prepareRequestBody(ctx, posts)
	fmt.Println(prompt)

	return nil, nil
}

func (g *Gemini) prepareRequestBody(_ context.Context, posts []*entity.Post) string {
	texts := ""
	for _, post := range posts {
		texts = fmt.Sprintf("%v\n\n%v- %v", texts, post.Id, post.Content)
	}

	return fmt.Sprintf(requestTokenTemplate, strings.TrimSpace(texts))
}
