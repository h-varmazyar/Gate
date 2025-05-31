package workers

import (
	"encoding/csv"
	"fmt"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/postProviders"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"os"
	"strconv"
	"strings"
)

type PostCollector interface {
	Collect(ctx context.Context, asset *chipmunkAPI.Asset, lastLoadedId string)
}

type CollectorRepository interface {
	SavePost(ctx context.Context, post *entity.Post) error
	BatchSave(ctx context.Context, posts []entity.Post) error
	OldestPost(ctx context.Context, provider chipmunkAPI.Provider) (*entity.Post, error)
}

type Collector struct {
	configs               Configs
	log                   *log.Logger
	sahamyabPostCollector PostCollector
	postCallbackChan      chan *entity.Post
	repository            CollectorRepository
}

func NewCollector(_ context.Context, log *log.Logger, configs Configs, repository CollectorRepository) *Collector {
	log.Infof("initializing collector worker")
	return &Collector{
		configs:          configs,
		log:              log,
		postCallbackChan: make(chan *entity.Post, 10000),
		repository:       repository,
	}
}

func (w *Collector) Start(ctx context.Context) {
	if !w.configs.Running {
		return
	}
	w.log.Infof("starting collector worker")
	networkConn := grpcext.NewConnection(w.configs.NetworkAddress)
	networkService := networkAPI.NewRequestServiceClient(networkConn)
	go w.listenPostCallback(ctx)
	go w.collectSahamyabPosts(ctx, networkService)
}

func (w *Collector) listenPostCallback(ctx context.Context) {
	if w.postCallbackChan == nil {
		err := errors.New(ctx, codes.FailedPrecondition).AddDetails("nil post callback channel not acceptable")
		w.log.WithError(err)
		return
	}

	//providersPosts := make(map[chipmunkAPI.Provider][]*entity.Post)
	// posts := make([]entity.Post, 0)
	for {
		select {
		case post := <-w.postCallbackChan:
			if err := w.repository.SavePost(ctx, post); err != nil {
				w.log.WithError(err).Error("failed to save post")
			}
			// posts = append(posts, *post)
			// if len(posts) > 9 {
			// 	w.log.Infof("saving batch posts at %v", time.Now())
			// 	if err := w.repository.BatchSave(ctx, posts); err != nil {
			// 		w.log.WithError(err).Errorf("failed to save post: %v", post)
			// 	}
			// 	posts = make([]entity.Post, 0)
			// }
			//provider := post.Provider
			////if post.PostedAt.Before(lastDay) {
			////	filename := lastDay.Format(time.DateOnly)
			////	if err := w.savePosts(filename, providersPosts[provider]); err != nil {
			////		w.log.WithError(err).Errorf("failed to save post of %v into %v", provider, filename)
			////	}
			////	lastDay.Add(-1 * time.Hour * 24)
			////	providersPosts[provider] = make([]*entity.Post, 0)
			////}
			//
			//if posts, ok := providersPosts[provider]; !ok {
			//	providersPosts[provider] = make([]*entity.Post, 0)
			//} else {
			//	providersPosts[provider] = append(providersPosts[provider], post)
			//
			//	if len(posts) == 1000 {
			//		filename := fmt.Sprintf("%v%v", strings.ToLower(provider.String()), count)
			//		if err := w.savePosts(filename, providersPosts[provider]); err != nil {
			//			w.log.WithError(err).Errorf("failed to save post of %v into %v", provider, filename)
			//		}
			//		providersPosts[provider] = make([]*entity.Post, 0)
			//		count++
			//	}
			//}

		}
	}
}

func (w *Collector) collectTwitterPosts(ctx context.Context) error {
	return errors.New(ctx, codes.Unimplemented)
}

func (w *Collector) collectSahamyabPosts(ctx context.Context, networkService networkAPI.RequestServiceClient) {
	w.log.Infof("initiating sahamyab collector")

	//oldestPost, err := w.repository.OldestPost(ctx, chipmunkAPI.Provider_SAHAMYAB)
	//if err != nil {
	//	if !sysErr.Is(err, gorm.ErrRecordNotFound) {
	//		w.log.WithError(err).Error("failed to find oldest post")
	//		return
	//	}
	//	w.log.Warnf("no oldest posts found")
	//}

	w.sahamyabPostCollector = postProviders.NewSahamyab(w.log, w.configs.SahamyabPostCollectorURL, networkService, w.postCallbackChan)

	w.sahamyabPostCollector.Collect(ctx, nil, "457934135")

}

func (w *Collector) collectCoinMarketCapPosts(ctx context.Context) error {
	return errors.New(ctx, codes.Unimplemented)
}

func (w *Collector) savePosts(filename string, posts []*entity.Post) error {
	w.log.Infof("saving %v", filename)
	if len(posts) == 0 {
		return nil
	}
	if !strings.HasSuffix(filename, ".csv") {
		filename += ".csv"
	}
	f, err := os.Create(filename)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return err
	}

	writer := csv.NewWriter(f)

	header := []string{
		"id",
		"postedAt",
		"content",
		"likeCount",
		"parentId",
		"senderUsername",
		"tags",
	}
	err = writer.Write(header)
	if err != nil {
		return err
	}

	for _, post := range posts {
		record := []string{
			post.Id,
			fmt.Sprint(post.PostedAt.Unix()),
			post.Content,
			fmt.Sprint(post.LikeCount),
			post.ParentId,
			post.SenderUsername,
			strings.Join(post.Tags, " "),
		}
		err = writer.Write(record)
		if err != nil {
			return err
		}
	}

	return nil
}

func findLastLoadedId(provider chipmunkAPI.Provider) (string, error) {
	info, err := os.ReadDir("./")
	if err != nil {
		return "", err
	}
	last := 0
	for _, entry := range info {
		fmt.Println(entry.Name())
		if strings.HasPrefix(entry.Name(), strings.ToLower(provider.String())) {
			found, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(entry.Name(), strings.ToLower(provider.String())), ".csv"))
			if err != nil {
				return "", err
			}

			if found > last {
				last = found
			}
		}
	}

	filename := fmt.Sprintf("%v%v.csv", strings.ToLower(provider.String()), last)
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return "", err
	}

	return records[len(records)-1][0], nil
}
