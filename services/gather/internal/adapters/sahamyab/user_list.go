package sahamyab

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/h-varmazyar/Gate/services/gather/internal/domain"
	"io"
	"net/http"
	"strconv"
	"time"
)

type post struct {
	ID             string `json:"id"`
	SendTime       string `json:"sendTime"`
	SenderName     string `json:"senderName"`
	SenderUsername string `json:"senderUsername"`
	Content        string `json:"content"`
	LikeCount      string `json:"likeCount"`
	CommentCount   string `json:"commentCount"`
	QuoteCount     string `json:"quoteCount"`
	RetwitCount    string `json:"retwitCount"`
	Type           string `json:"type"`
	ScoredPostDate string `json:"scoredPostDate"`
	ParentID       string `json:"parentId"`
}

type posts struct {
	baseResponse
	Items []post `json:"items"`
}

func (s *Sahamyab) GetUserPageList(_ context.Context, input domain.GetScoredSahamyabPost) (domain.SahamyabPostList, error) {
	params := map[string]any{
		"page": input.Page,
	}

	if input.ScoredPostDate != "" {
		params["scoredPostDate"] = input.ScoredPostDate
	}

	jsonData, err := json.Marshal(params)
	if err != nil {
		return domain.SahamyabPostList{}, err
	}
	req, err := http.NewRequest(http.MethodPost, "https://www.sahamyab.com/guest/twiter/userPageList?v=0.1", bytes.NewBuffer(jsonData))
	if err != nil {
		return domain.SahamyabPostList{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return domain.SahamyabPostList{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.SahamyabPostList{}, err
	}

	respEntity := posts{}
	err = json.Unmarshal(respBody, &respEntity)
	if err != nil {
		return domain.SahamyabPostList{}, err
	}

	if !respEntity.Success {
		return domain.SahamyabPostList{}, fmt.Errorf("invalid response: %v", respEntity.ErrorTitle)
	}

	posts := make([]*domain.SahamyabPost, 0)
	for _, item := range respEntity.Items {
		post, err := parsePost(item)
		if err != nil {
			continue
		}
		posts = append(posts, &post)
	}

	postList := domain.SahamyabPostList{
		Items: posts,
	}
	fmt.Println(input.ScoredPostDate, postList.Items[len(postList.Items)-1].ScoredPostDate)

	return postList, nil
}

func parsePost(item post) (domain.SahamyabPost, error) {
	post := domain.SahamyabPost{
		SenderName:     item.SenderName,
		SenderUsername: item.SenderUsername,
		Content:        item.Content,
		Type:           item.Type,
		ScoredPostDate: item.ScoredPostDate,
	}

	var err error
	post.SendTime, err = time.Parse(time.RFC3339, item.SendTime)
	if err != nil {
		return domain.SahamyabPost{}, err
	}

	post.ID, err = strconv.ParseInt(item.ID, 10, 64)
	if err != nil {
		return domain.SahamyabPost{}, err
	}

	if item.LikeCount != "" {
		post.LikeCount, _ = strconv.Atoi(item.LikeCount)
	}

	if item.RetwitCount != "" {
		post.RetwitCount, _ = strconv.Atoi(item.RetwitCount)
	}

	if item.CommentCount != "" {
		post.CommentCount, _ = strconv.Atoi(item.CommentCount)
	}

	if item.QuoteCount != "" {
		post.QuoteCount, _ = strconv.Atoi(item.QuoteCount)
	}

	if item.ParentID != "" {
		post.ParentID, _ = strconv.ParseInt(item.ParentID, 10, 64)
	}

	return post, nil
}
