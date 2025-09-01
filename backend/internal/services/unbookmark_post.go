package services

import (
	"context"
)

type UnbookmarkPost struct{}

type UnbookmarkPostReq struct {
	Slug   string `json:"-"`
	UserID string `json:"-"`
}

type UnbookmarkPostResp struct {
	Bookmarked bool `json:"bookmarked"`
}

func NewUnbookmarkPost() *UnbookmarkPost {
	return &UnbookmarkPost{}
}

func (s *UnbookmarkPost) Exec(ctx context.Context, req *UnbookmarkPostReq) (*UnbookmarkPostResp, error) {
	return &UnbookmarkPostResp{
		Bookmarked: false,
	}, nil
}