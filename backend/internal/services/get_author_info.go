package services

import (
	"context"
)

type GetAuthorInfo struct{}

type GetAuthorInfoReq struct {
	AuthorID string
}

type TopPostInfo struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	LikesCount int    `json:"likes_count"`
}

type GetAuthorInfoResp struct {
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	PostsCount     int           `json:"posts_count"`
	FollowersCount int           `json:"followers_count"`
	TopPosts       []TopPostInfo `json:"top_posts"`
}

func NewGetAuthorInfo() *GetAuthorInfo {
	return &GetAuthorInfo{}
}

func (s *GetAuthorInfo) Exec(ctx context.Context, req *GetAuthorInfoReq) (*GetAuthorInfoResp, error) {
	return &GetAuthorInfoResp{
		ID:             req.AuthorID,
		Name:           "Sample Author",
		PostsCount:     12,
		FollowersCount: 128,
		TopPosts: []TopPostInfo{
			{ID: "p1", Title: "Post A", Slug: "post-a", LikesCount: 420},
			{ID: "p2", Title: "Post B", Slug: "post-b", LikesCount: 120},
		},
	}, nil
}