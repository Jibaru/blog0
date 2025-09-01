package services

import (
	"context"
)

type ToggleLike struct{}

type ToggleLikeReq struct {
	Slug   string `json:"-"`
	UserID string `json:"-"`
}

type ToggleLikeResp struct {
	Liked      bool `json:"liked"`
	LikesCount int  `json:"likes_count"`
}

func NewToggleLike() *ToggleLike {
	return &ToggleLike{}
}

func (s *ToggleLike) Exec(ctx context.Context, req *ToggleLikeReq) (*ToggleLikeResp, error) {
	return &ToggleLikeResp{
		Liked:      true,
		LikesCount: 43,
	}, nil
}