package services

import (
	"context"
	"time"
)

type GetPostBySlug struct{}

type GetPostBySlugReq struct {
	Slug string
}

type AuthorInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CommentInfo struct {
	ID       string     `json:"id"`
	Author   AuthorInfo `json:"author"`
	ParentID *string    `json:"parent_id"`
	Body     string     `json:"body"`
	CreateAt time.Time  `json:"created_at"`
}

type GetPostBySlugResp struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Slug        string        `json:"slug"`
	RawMarkdown string        `json:"raw_markdown"`
	Author      AuthorInfo    `json:"author"`
	PublishedAt time.Time     `json:"published_at"`
	LikesCount  int           `json:"likes_count"`
	Comments    []CommentInfo `json:"comments"`
}

func NewGetPostBySlug() *GetPostBySlug {
	return &GetPostBySlug{}
}

func (s *GetPostBySlug) Exec(ctx context.Context, req *GetPostBySlugReq) (*GetPostBySlugResp, error) {
	return &GetPostBySlugResp{
		ID:          "uuid-del-post",
		Title:       "Sample Post",
		Slug:        req.Slug,
		RawMarkdown: "## Sample Content\nThis is a sample post.",
		Author:      AuthorInfo{ID: "uuid-author", Name: "Sample Author"},
		PublishedAt: time.Now(),
		LikesCount:  0,
		Comments:    []CommentInfo{},
	}, nil
}