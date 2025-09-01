package services

import (
	"context"
)

type FollowUser struct{}

type FollowUserReq struct {
	AuthorID string `json:"-"`
	UserID   string `json:"-"`
}

type FollowUserResp struct {
	Following      bool `json:"following"`
	FollowersCount int  `json:"followers_count"`
}

func NewFollowUser() *FollowUser {
	return &FollowUser{}
}

func (s *FollowUser) Exec(ctx context.Context, req *FollowUserReq) (*FollowUserResp, error) {
	return &FollowUserResp{
		Following:      true,
		FollowersCount: 128,
	}, nil
}