package services

import (
	"context"
	"time"
)

type BookmarkPost struct{}

type BookmarkPostReq struct {
	Slug   string `json:"-"`
	UserID string `json:"-"`
}

type BookmarkPostResp struct {
	Bookmarked bool      `json:"bookmarked"`
	BookmarkID string    `json:"bookmark_id"`
	PostSlug   string    `json:"post_slug"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewBookmarkPost() *BookmarkPost {
	return &BookmarkPost{}
}

func (s *BookmarkPost) Exec(ctx context.Context, req *BookmarkPostReq) (*BookmarkPostResp, error) {
	return &BookmarkPostResp{
		Bookmarked: true,
		BookmarkID: "uuid-bookmark",
		PostSlug:   req.Slug,
		CreatedAt:  time.Now(),
	}, nil
}