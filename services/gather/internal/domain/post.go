package domain

import "time"

type GetScoredSahamyabPost struct {
	Page           int    `json:"page"`
	ScoredPostDate string `json:"scoredPostDate"`
}
type SahamyabPost struct {
	ID             int64     `json:"id"`
	SendTime       time.Time `json:"sendTime"`
	SenderName     string    `json:"senderName"`
	SenderUsername string    `json:"senderUsername"`
	Content        string    `json:"content"`
	LikeCount      int       `json:"likeCount"`
	RetwitCount    int       `json:"retwitCount"`
	CommentCount   int       `json:"commentCount"`
	QuoteCount     int       `json:"quoteCount"`
	Type           string    `json:"type"`
	ScoredPostDate string    `json:"scoredPostDate"`
	ParentID       int64     `json:"parentId"`
}
type SahamyabPostList struct {
	Items []*SahamyabPost `json:"items"`
}
