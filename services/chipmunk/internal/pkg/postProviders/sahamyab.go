package postProviders

import (
	"encoding/json"
	"fmt"
	"github.com/h-varmazyar/Gate/pkg/errors"
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Sahamyab struct {
	PostsCollectorURL      string
	SinglePostCollectorURL string
	NetworkService         networkAPI.RequestServiceClient
	PostCallbackChan       chan *entity.Post
	Log                    *log.Logger
	page                   int64
	//lastId                 string
}

type PostResp struct {
	ErrorCode  string  `json:"errorCode"`
	ErrorTitle string  `json:"errorTitle"`
	HasMore    bool    `json:"hasMore"`
	Items      []*Item `json:"items"`
	Success    bool    `json:"success"`
	*Item
}

type Item struct {
	Id             string `json:"id"`
	Content        string `json:"content"`
	SendTime       string `json:"sendTime"`
	SenderUsername string `json:"senderUsername"`
	Type           string `json:"type"`
	Advertise      bool   `json:"advertise"`
	ParentId       string `json:"parentId"`
	LikeCount      string `json:"likeCount"`
}

func NewSahamyab(log *log.Logger, postCollectorURL string, networkService networkAPI.RequestServiceClient, postCallbackChan chan *entity.Post) *Sahamyab {
	return &Sahamyab{
		PostsCollectorURL: postCollectorURL,
		// SinglePostCollectorURL: "https://www.sahamyab.com/guest/twiter/item?v=0.1",
		NetworkService:   networkService,
		PostCallbackChan: postCallbackChan,
		Log:              log,
	}
}

func (c *Sahamyab) Collect(_ context.Context, _ *chipmunkAPI.Asset, lastLoadedId string) {
	c.Log.Infof("collecting sahamyab posts. last id is %v", lastLoadedId)
	c.page = int64(12)

	lastId, err := strconv.Atoi(lastLoadedId)
	if err != nil {
		return
	}

	ticker := time.NewTicker(time.Minute / 2)
	for {
		select {
		case <-ticker.C:
			if time.Now().Hour() == 8 {
				c.Log.Infof("sleeping in market time...")
				time.Sleep(time.Hour * 7)
				c.Log.Infof("resuming after market time...")
			}
			go c.sendRequest(context.Background(), lastId)
		}
		lastId++
	}
}

func (c *Sahamyab) sendRequest(ctx context.Context, lastId int) {
	req := &networkAPI.Request{
		Method:    networkAPI.Request_POST,
		Headers:   c.prepareHeaders(),
		Params:    c.prepareParams(fmt.Sprintf("%v", lastId)),
		IssueTime: time.Now().Unix(),
	}
	if c.SinglePostCollectorURL != "" {
		req.Endpoint = c.SinglePostCollectorURL
	} else {
		req.Endpoint = c.PostsCollectorURL
	}
	resp, err := c.NetworkService.Do(ctx, req)
	if err != nil {
		c.handleErr(ctx, err)
		return
	}
	if resp.GetCode() == http.StatusOK {
		posts, err := c.parseResponseBody(ctx, resp.GetBody())
		if err != nil || len(posts) == 0 {
			return
		}
		err = c.publishPosts(ctx, posts)
		if err != nil {
			c.Log.WithError(err).Errorf("failed to publish sahamyab post")
			return
		}
	} else {
		c.handleFailedStatus(ctx, resp.GetCode())
	}
}

func (c *Sahamyab) prepareHeaders() []*networkAPI.KV {
	return []*networkAPI.KV{
		{
			Key:   "Accept",
			Value: &networkAPI.KV_String_{String_: "application/json, text/plain, */*"},
		},
		//{
		//	Key:   "Accept-Encoding",
		//	Value: &networkAPI.KV_String_{String_: "gzip, deflate, br"},
		//},
		//{
		//	Key:   "Accept-Language",
		//	Value: &networkAPI.KV_String_{String_: "en-US,en;q=0.9"},
		//},
		//{
		//	Key:   "Content-Type",
		//	Value: &networkAPI.KV_String_{String_: "application/json"},
		//},
		//{
		//	Key:   "Origin",
		//	Value: &networkAPI.KV_String_{String_: "https://www.sahamyab.com"},
		//},
		//{
		//	Key:   "User-Agent",
		//	Value: &networkAPI.KV_String_{String_: "Mozilla/6.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/637.36"},
		//},
		//{
		//	Key:   "Referer",
		//	Value: &networkAPI.KV_String_{String_: "https://www.sahamyab.com/stocktwits"},
		//},
		//{
		//	Key:   "Sec-Ch-Ua",
		//	Value: &networkAPI.KV_String_{String_: `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`},
		//},
		//{
		//	Key:   "Sec-Ch-Ua-Mobile",
		//	Value: &networkAPI.KV_String_{String_: "?0"},
		//},
		//{
		//	Key:   "Sec-Ch-Ua-Platform",
		//	Value: &networkAPI.KV_String_{String_: "Linux"},
		//},
		//{
		//	Key:   "Cookie",
		//	Value: &networkAPI.KV_String_{String_: "_ga_EF12XEFLM6=GS1.1.1683704333.1.0.1683704333.0.0.0; _gid=GA1.2.577816098.1707575453; _gat_UA-39858392-1=1; _ga=GA1.2.1926163846.1686639997; _ga_15TLTHGBVQ=GS1.1.1707575452.57.1.1707575506.0.0.0"},
		//},
		//{
		//	Key:   "",
		//	Value: &networkAPI.KV_String_{String_: ""},
		//},
		//{
		//	Key:   "",
		//	Value: &networkAPI.KV_String_{String_: ""},
		//},
	}
}

func (c *Sahamyab) prepareParams(lastId string) []*networkAPI.KV {
	params := make([]*networkAPI.KV, 0)

	pageKV := &networkAPI.KV{
		Key:   "page",
		Value: &networkAPI.KV_Integer{Integer: c.page},
	}
	params = append(params, pageKV)
	if lastId != "" {
		lastIdKV := &networkAPI.KV{
			Key:   "id",
			Value: &networkAPI.KV_String_{String_: lastId},
		}
		params = append(params, lastIdKV)
	}

	return params
}

func (c *Sahamyab) parseResponseBody(ctx context.Context, body string) ([]*entity.Post, error) {
	resp := new(PostResp)
	if err := json.Unmarshal([]byte(body), resp); err != nil {
		c.Log.WithError(err).Errorf("failed to parse response")
		return nil, err
	}

	if !resp.Success {
		return nil, errors.New(ctx, codes.Unknown).AddDetails("unsuccessful data gathering")
	}

	posts := make([]*entity.Post, 0)

	if len(resp.Items) > 0 {
		for _, item := range resp.Items {
			if item.Advertise {
				continue
			}

			if post := c.parsePost(item); post != nil {
				posts = append(posts, post)
			}
		}
	} else {
		posts = append(posts, c.parsePost(resp.Item))
	}

	return posts, nil
}

func (c *Sahamyab) parsePost(item *Item) *entity.Post {
	if !hasValidContent(item.Content) {
		return nil
	}

	sendTime, err := time.Parse(time.RFC3339, item.SendTime)
	if err != nil {
		c.Log.Infof("failed to parse time: %v", err)
		return nil
	}

	likeCount, _ := strconv.Atoi(item.LikeCount)

	return &entity.Post{
		Id:             item.Id,
		PostedAt:       sendTime,
		Content:        item.Content,
		LikeCount:      uint32(likeCount),
		ParentId:       item.ParentId,
		SenderUsername: item.SenderUsername,
		Provider:       chipmunkAPI.Provider_SAHAMYAB,
		Tags:           fetchTags(item.Content),
	}
}

func (c *Sahamyab) publishPosts(ctx context.Context, posts []*entity.Post) error {
	if c.PostCallbackChan == nil {
		return errors.New(ctx, codes.FailedPrecondition).AddDetails("nil post callback channel for sahamyab")
	}
	for _, post := range posts {
		if post != nil {
			c.PostCallbackChan <- post
		}
	}
	return nil
}

func (c *Sahamyab) handleErr(_ context.Context, err error) {
	c.Log.WithError(err).Errorf("sahamyab do request error")
}

func (c *Sahamyab) handleFailedStatus(_ context.Context, code int32) {
	c.Log.Errorf("sahamyab status code error: %v", code)
	switch code {
	case 429:
		c.page++
		if c.page == 13 {
			c.page = 0
			time.Sleep(time.Hour * 3)
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
